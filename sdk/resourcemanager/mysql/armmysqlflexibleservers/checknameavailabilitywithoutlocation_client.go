//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmysqlflexibleservers

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// CheckNameAvailabilityWithoutLocationClient contains the methods for the CheckNameAvailabilityWithoutLocation group.
// Don't use this type directly, use NewCheckNameAvailabilityWithoutLocationClient() instead.
type CheckNameAvailabilityWithoutLocationClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewCheckNameAvailabilityWithoutLocationClient creates a new instance of CheckNameAvailabilityWithoutLocationClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewCheckNameAvailabilityWithoutLocationClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*CheckNameAvailabilityWithoutLocationClient, error) {
	cl, err := arm.NewClient(moduleName+".CheckNameAvailabilityWithoutLocationClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &CheckNameAvailabilityWithoutLocationClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Execute - Check the availability of name for server
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01-preview
//   - nameAvailabilityRequest - The required parameters for checking if server name is available.
//   - options - CheckNameAvailabilityWithoutLocationClientExecuteOptions contains the optional parameters for the CheckNameAvailabilityWithoutLocationClient.Execute
//     method.
func (client *CheckNameAvailabilityWithoutLocationClient) Execute(ctx context.Context, nameAvailabilityRequest NameAvailabilityRequest, options *CheckNameAvailabilityWithoutLocationClientExecuteOptions) (CheckNameAvailabilityWithoutLocationClientExecuteResponse, error) {
	var err error
	req, err := client.executeCreateRequest(ctx, nameAvailabilityRequest, options)
	if err != nil {
		return CheckNameAvailabilityWithoutLocationClientExecuteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return CheckNameAvailabilityWithoutLocationClientExecuteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return CheckNameAvailabilityWithoutLocationClientExecuteResponse{}, err
	}
	resp, err := client.executeHandleResponse(httpResp)
	return resp, err
}

// executeCreateRequest creates the Execute request.
func (client *CheckNameAvailabilityWithoutLocationClient) executeCreateRequest(ctx context.Context, nameAvailabilityRequest NameAvailabilityRequest, options *CheckNameAvailabilityWithoutLocationClientExecuteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.DBforMySQL/checkNameAvailability"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, nameAvailabilityRequest); err != nil {
		return nil, err
	}
	return req, nil
}

// executeHandleResponse handles the Execute response.
func (client *CheckNameAvailabilityWithoutLocationClient) executeHandleResponse(resp *http.Response) (CheckNameAvailabilityWithoutLocationClientExecuteResponse, error) {
	result := CheckNameAvailabilityWithoutLocationClientExecuteResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NameAvailability); err != nil {
		return CheckNameAvailabilityWithoutLocationClientExecuteResponse{}, err
	}
	return result, nil
}
