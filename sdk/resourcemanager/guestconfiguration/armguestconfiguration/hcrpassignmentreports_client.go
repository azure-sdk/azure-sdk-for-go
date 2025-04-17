// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armguestconfiguration

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

// HCRPAssignmentReportsClient contains the methods for the GuestConfigurationHCRPAssignmentReports group.
// Don't use this type directly, use NewHCRPAssignmentReportsClient() instead.
type HCRPAssignmentReportsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewHCRPAssignmentReportsClient creates a new instance of HCRPAssignmentReportsClient with the specified values.
//   - subscriptionID - Subscription ID which uniquely identify Microsoft Azure subscription. The subscription ID forms part of
//     the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewHCRPAssignmentReportsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*HCRPAssignmentReportsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &HCRPAssignmentReportsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Get a report for the guest configuration assignment, by reportId.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-05
//   - resourceGroupName - The resource group name.
//   - guestConfigurationAssignmentName - The guest configuration assignment name.
//   - reportID - The GUID for the guest configuration assignment report.
//   - machineName - The name of the ARC machine.
//   - options - HCRPAssignmentReportsClientGetOptions contains the optional parameters for the HCRPAssignmentReportsClient.Get
//     method.
func (client *HCRPAssignmentReportsClient) Get(ctx context.Context, resourceGroupName string, guestConfigurationAssignmentName string, reportID string, machineName string, options *HCRPAssignmentReportsClientGetOptions) (HCRPAssignmentReportsClientGetResponse, error) {
	var err error
	const operationName = "HCRPAssignmentReportsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, guestConfigurationAssignmentName, reportID, machineName, options)
	if err != nil {
		return HCRPAssignmentReportsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HCRPAssignmentReportsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return HCRPAssignmentReportsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *HCRPAssignmentReportsClient) getCreateRequest(ctx context.Context, resourceGroupName string, guestConfigurationAssignmentName string, reportID string, machineName string, _ *HCRPAssignmentReportsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridCompute/machines/{machineName}/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments/{guestConfigurationAssignmentName}/reports/{reportId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if guestConfigurationAssignmentName == "" {
		return nil, errors.New("parameter guestConfigurationAssignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{guestConfigurationAssignmentName}", url.PathEscape(guestConfigurationAssignmentName))
	if reportID == "" {
		return nil, errors.New("parameter reportID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{reportId}", url.PathEscape(reportID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if machineName == "" {
		return nil, errors.New("parameter machineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{machineName}", url.PathEscape(machineName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-05")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *HCRPAssignmentReportsClient) getHandleResponse(resp *http.Response) (HCRPAssignmentReportsClientGetResponse, error) {
	result := HCRPAssignmentReportsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AssignmentReport); err != nil {
		return HCRPAssignmentReportsClientGetResponse{}, err
	}
	return result, nil
}

// List - List all reports for the guest configuration assignment, latest report first.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-05
//   - resourceGroupName - The resource group name.
//   - guestConfigurationAssignmentName - The guest configuration assignment name.
//   - machineName - The name of the ARC machine.
//   - options - HCRPAssignmentReportsClientListOptions contains the optional parameters for the HCRPAssignmentReportsClient.List
//     method.
func (client *HCRPAssignmentReportsClient) List(ctx context.Context, resourceGroupName string, guestConfigurationAssignmentName string, machineName string, options *HCRPAssignmentReportsClientListOptions) (HCRPAssignmentReportsClientListResponse, error) {
	var err error
	const operationName = "HCRPAssignmentReportsClient.List"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listCreateRequest(ctx, resourceGroupName, guestConfigurationAssignmentName, machineName, options)
	if err != nil {
		return HCRPAssignmentReportsClientListResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HCRPAssignmentReportsClientListResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return HCRPAssignmentReportsClientListResponse{}, err
	}
	resp, err := client.listHandleResponse(httpResp)
	return resp, err
}

// listCreateRequest creates the List request.
func (client *HCRPAssignmentReportsClient) listCreateRequest(ctx context.Context, resourceGroupName string, guestConfigurationAssignmentName string, machineName string, _ *HCRPAssignmentReportsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridCompute/machines/{machineName}/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments/{guestConfigurationAssignmentName}/reports"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if guestConfigurationAssignmentName == "" {
		return nil, errors.New("parameter guestConfigurationAssignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{guestConfigurationAssignmentName}", url.PathEscape(guestConfigurationAssignmentName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if machineName == "" {
		return nil, errors.New("parameter machineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{machineName}", url.PathEscape(machineName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-05")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *HCRPAssignmentReportsClient) listHandleResponse(resp *http.Response) (HCRPAssignmentReportsClientListResponse, error) {
	result := HCRPAssignmentReportsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AssignmentReportList); err != nil {
		return HCRPAssignmentReportsClientListResponse{}, err
	}
	return result, nil
}
