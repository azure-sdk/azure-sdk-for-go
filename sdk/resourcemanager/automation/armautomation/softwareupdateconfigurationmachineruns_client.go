//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armautomation

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

// SoftwareUpdateConfigurationMachineRunsClient contains the methods for the SoftwareUpdateConfigurationMachineRuns group.
// Don't use this type directly, use NewSoftwareUpdateConfigurationMachineRunsClient() instead.
type SoftwareUpdateConfigurationMachineRunsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewSoftwareUpdateConfigurationMachineRunsClient creates a new instance of SoftwareUpdateConfigurationMachineRunsClient with the specified values.
//   - subscriptionID - Gets subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID
//     forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewSoftwareUpdateConfigurationMachineRunsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SoftwareUpdateConfigurationMachineRunsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &SoftwareUpdateConfigurationMachineRunsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// GetByID - Get a single software update configuration machine run by Id.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-01
//   - resourceGroupName - Name of an Azure Resource group.
//   - automationAccountName - The name of the automation account.
//   - softwareUpdateConfigurationMachineRunID - The Id of the software update configuration machine run.
//   - options - SoftwareUpdateConfigurationMachineRunsClientGetByIDOptions contains the optional parameters for the SoftwareUpdateConfigurationMachineRunsClient.GetByID
//     method.
func (client *SoftwareUpdateConfigurationMachineRunsClient) GetByID(ctx context.Context, resourceGroupName string, automationAccountName string, softwareUpdateConfigurationMachineRunID string, options *SoftwareUpdateConfigurationMachineRunsClientGetByIDOptions) (SoftwareUpdateConfigurationMachineRunsClientGetByIDResponse, error) {
	var err error
	const operationName = "SoftwareUpdateConfigurationMachineRunsClient.GetByID"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getByIDCreateRequest(ctx, resourceGroupName, automationAccountName, softwareUpdateConfigurationMachineRunID, options)
	if err != nil {
		return SoftwareUpdateConfigurationMachineRunsClientGetByIDResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SoftwareUpdateConfigurationMachineRunsClientGetByIDResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SoftwareUpdateConfigurationMachineRunsClientGetByIDResponse{}, err
	}
	resp, err := client.getByIDHandleResponse(httpResp)
	return resp, err
}

// getByIDCreateRequest creates the GetByID request.
func (client *SoftwareUpdateConfigurationMachineRunsClient) getByIDCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, softwareUpdateConfigurationMachineRunID string, options *SoftwareUpdateConfigurationMachineRunsClientGetByIDOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/softwareUpdateConfigurationMachineRuns/{softwareUpdateConfigurationMachineRunId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if softwareUpdateConfigurationMachineRunID == "" {
		return nil, errors.New("parameter softwareUpdateConfigurationMachineRunID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{softwareUpdateConfigurationMachineRunId}", url.PathEscape(softwareUpdateConfigurationMachineRunID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.ClientRequestID != nil {
		req.Raw().Header["clientRequestId"] = []string{*options.ClientRequestID}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getByIDHandleResponse handles the GetByID response.
func (client *SoftwareUpdateConfigurationMachineRunsClient) getByIDHandleResponse(resp *http.Response) (SoftwareUpdateConfigurationMachineRunsClientGetByIDResponse, error) {
	result := SoftwareUpdateConfigurationMachineRunsClientGetByIDResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SoftwareUpdateConfigurationMachineRun); err != nil {
		return SoftwareUpdateConfigurationMachineRunsClientGetByIDResponse{}, err
	}
	return result, nil
}

// List - Return list of software update configuration machine runs
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-01
//   - resourceGroupName - Name of an Azure Resource group.
//   - automationAccountName - The name of the automation account.
//   - options - SoftwareUpdateConfigurationMachineRunsClientListOptions contains the optional parameters for the SoftwareUpdateConfigurationMachineRunsClient.List
//     method.
func (client *SoftwareUpdateConfigurationMachineRunsClient) List(ctx context.Context, resourceGroupName string, automationAccountName string, options *SoftwareUpdateConfigurationMachineRunsClientListOptions) (SoftwareUpdateConfigurationMachineRunsClientListResponse, error) {
	var err error
	const operationName = "SoftwareUpdateConfigurationMachineRunsClient.List"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listCreateRequest(ctx, resourceGroupName, automationAccountName, options)
	if err != nil {
		return SoftwareUpdateConfigurationMachineRunsClientListResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SoftwareUpdateConfigurationMachineRunsClientListResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SoftwareUpdateConfigurationMachineRunsClientListResponse{}, err
	}
	resp, err := client.listHandleResponse(httpResp)
	return resp, err
}

// listCreateRequest creates the List request.
func (client *SoftwareUpdateConfigurationMachineRunsClient) listCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, options *SoftwareUpdateConfigurationMachineRunsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/softwareUpdateConfigurationMachineRuns"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-11-01")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", *options.Skip)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", *options.Top)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.ClientRequestID != nil {
		req.Raw().Header["clientRequestId"] = []string{*options.ClientRequestID}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *SoftwareUpdateConfigurationMachineRunsClient) listHandleResponse(resp *http.Response) (SoftwareUpdateConfigurationMachineRunsClientListResponse, error) {
	result := SoftwareUpdateConfigurationMachineRunsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SoftwareUpdateConfigurationMachineRunListResult); err != nil {
		return SoftwareUpdateConfigurationMachineRunsClientListResponse{}, err
	}
	return result, nil
}
