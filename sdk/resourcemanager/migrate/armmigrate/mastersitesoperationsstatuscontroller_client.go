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

// MasterSitesOperationsStatusControllerClient contains the methods for the MasterSitesOperationsStatusController group.
// Don't use this type directly, use NewMasterSitesOperationsStatusControllerClient() instead.
type MasterSitesOperationsStatusControllerClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewMasterSitesOperationsStatusControllerClient creates a new instance of MasterSitesOperationsStatusControllerClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewMasterSitesOperationsStatusControllerClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*MasterSitesOperationsStatusControllerClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &MasterSitesOperationsStatusControllerClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// GetVmwareOperationStatus - A operation status resource belonging to a master site resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteName - Site name
//   - operationStatusName - Operation status Arm Name.
//   - options - MasterSitesOperationsStatusControllerClientGetVmwareOperationStatusOptions contains the optional parameters for
//     the MasterSitesOperationsStatusControllerClient.GetVmwareOperationStatus method.
func (client *MasterSitesOperationsStatusControllerClient) GetVmwareOperationStatus(ctx context.Context, resourceGroupName string, siteName string, operationStatusName string, options *MasterSitesOperationsStatusControllerClientGetVmwareOperationStatusOptions) (MasterSitesOperationsStatusControllerClientGetVmwareOperationStatusResponse, error) {
	var err error
	const operationName = "MasterSitesOperationsStatusControllerClient.GetVmwareOperationStatus"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getVmwareOperationStatusCreateRequest(ctx, resourceGroupName, siteName, operationStatusName, options)
	if err != nil {
		return MasterSitesOperationsStatusControllerClientGetVmwareOperationStatusResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MasterSitesOperationsStatusControllerClientGetVmwareOperationStatusResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return MasterSitesOperationsStatusControllerClientGetVmwareOperationStatusResponse{}, err
	}
	resp, err := client.getVmwareOperationStatusHandleResponse(httpResp)
	return resp, err
}

// getVmwareOperationStatusCreateRequest creates the GetVmwareOperationStatus request.
func (client *MasterSitesOperationsStatusControllerClient) getVmwareOperationStatusCreateRequest(ctx context.Context, resourceGroupName string, siteName string, operationStatusName string, options *MasterSitesOperationsStatusControllerClientGetVmwareOperationStatusOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/masterSites/{siteName}/operationsStatus/{operationStatusName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if siteName == "" {
		return nil, errors.New("parameter siteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{siteName}", url.PathEscape(siteName))
	if operationStatusName == "" {
		return nil, errors.New("parameter operationStatusName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{operationStatusName}", url.PathEscape(operationStatusName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getVmwareOperationStatusHandleResponse handles the GetVmwareOperationStatus response.
func (client *MasterSitesOperationsStatusControllerClient) getVmwareOperationStatusHandleResponse(resp *http.Response) (MasterSitesOperationsStatusControllerClientGetVmwareOperationStatusResponse, error) {
	result := MasterSitesOperationsStatusControllerClientGetVmwareOperationStatusResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OperationStatus); err != nil {
		return MasterSitesOperationsStatusControllerClientGetVmwareOperationStatusResponse{}, err
	}
	return result, nil
}