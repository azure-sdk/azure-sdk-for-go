//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armpostgresqlflexibleservers

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// CheckNameAvailabilityWithLocationClient contains the methods for the CheckNameAvailabilityWithLocation group.
// Don't use this type directly, use NewCheckNameAvailabilityWithLocationClient() instead.
type CheckNameAvailabilityWithLocationClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewCheckNameAvailabilityWithLocationClient creates a new instance of CheckNameAvailabilityWithLocationClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewCheckNameAvailabilityWithLocationClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*CheckNameAvailabilityWithLocationClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &CheckNameAvailabilityWithLocationClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Execute - Check the availability of name for resource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-12-01
//   - locationName - The name of the location.
//   - nameAvailabilityRequest - The required parameters for checking if resource name is available.
//   - options - CheckNameAvailabilityWithLocationClientExecuteOptions contains the optional parameters for the CheckNameAvailabilityWithLocationClient.Execute
//     method.
func (client *CheckNameAvailabilityWithLocationClient) Execute(ctx context.Context, locationName string, nameAvailabilityRequest CheckNameAvailabilityRequest, options *CheckNameAvailabilityWithLocationClientExecuteOptions) (CheckNameAvailabilityWithLocationClientExecuteResponse, error) {
	req, err := client.executeCreateRequest(ctx, locationName, nameAvailabilityRequest, options)
	if err != nil {
		return CheckNameAvailabilityWithLocationClientExecuteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CheckNameAvailabilityWithLocationClientExecuteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return CheckNameAvailabilityWithLocationClientExecuteResponse{}, runtime.NewResponseError(resp)
	}
	return client.executeHandleResponse(resp)
}

// executeCreateRequest creates the Execute request.
func (client *CheckNameAvailabilityWithLocationClient) executeCreateRequest(ctx context.Context, locationName string, nameAvailabilityRequest CheckNameAvailabilityRequest, options *CheckNameAvailabilityWithLocationClientExecuteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.DBforPostgreSQL/locations/{locationName}/checkNameAvailability"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if locationName == "" {
		return nil, errors.New("parameter locationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{locationName}", url.PathEscape(locationName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, nameAvailabilityRequest)
}

// executeHandleResponse handles the Execute response.
func (client *CheckNameAvailabilityWithLocationClient) executeHandleResponse(resp *http.Response) (CheckNameAvailabilityWithLocationClientExecuteResponse, error) {
	result := CheckNameAvailabilityWithLocationClientExecuteResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NameAvailability); err != nil {
		return CheckNameAvailabilityWithLocationClientExecuteResponse{}, err
	}
	return result, nil
}
