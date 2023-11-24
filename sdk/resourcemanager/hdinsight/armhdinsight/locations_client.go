//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhdinsight

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

// LocationsClient contains the methods for the Locations group.
// Don't use this type directly, use NewLocationsClient() instead.
type LocationsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewLocationsClient creates a new instance of LocationsClient with the specified values.
//   - subscriptionID - The subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID
//     forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewLocationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*LocationsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &LocationsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CheckNameAvailability - Check the cluster name is available or not.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-08-15-preview
//   - location - The Azure location (region) for which to make the request.
//   - options - LocationsClientCheckNameAvailabilityOptions contains the optional parameters for the LocationsClient.CheckNameAvailability
//     method.
func (client *LocationsClient) CheckNameAvailability(ctx context.Context, location string, parameters NameAvailabilityCheckRequestParameters, options *LocationsClientCheckNameAvailabilityOptions) (LocationsClientCheckNameAvailabilityResponse, error) {
	var err error
	const operationName = "LocationsClient.CheckNameAvailability"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.checkNameAvailabilityCreateRequest(ctx, location, parameters, options)
	if err != nil {
		return LocationsClientCheckNameAvailabilityResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return LocationsClientCheckNameAvailabilityResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return LocationsClientCheckNameAvailabilityResponse{}, err
	}
	resp, err := client.checkNameAvailabilityHandleResponse(httpResp)
	return resp, err
}

// checkNameAvailabilityCreateRequest creates the CheckNameAvailability request.
func (client *LocationsClient) checkNameAvailabilityCreateRequest(ctx context.Context, location string, parameters NameAvailabilityCheckRequestParameters, options *LocationsClientCheckNameAvailabilityOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.HDInsight/locations/{location}/checkNameAvailability"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-08-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// checkNameAvailabilityHandleResponse handles the CheckNameAvailability response.
func (client *LocationsClient) checkNameAvailabilityHandleResponse(resp *http.Response) (LocationsClientCheckNameAvailabilityResponse, error) {
	result := LocationsClientCheckNameAvailabilityResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NameAvailabilityCheckResult); err != nil {
		return LocationsClientCheckNameAvailabilityResponse{}, err
	}
	return result, nil
}

// GetAzureAsyncOperationStatus - Get the async operation status.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-08-15-preview
//   - location - The Azure location (region) for which to make the request.
//   - operationID - The long running operation id.
//   - options - LocationsClientGetAzureAsyncOperationStatusOptions contains the optional parameters for the LocationsClient.GetAzureAsyncOperationStatus
//     method.
func (client *LocationsClient) GetAzureAsyncOperationStatus(ctx context.Context, location string, operationID string, options *LocationsClientGetAzureAsyncOperationStatusOptions) (LocationsClientGetAzureAsyncOperationStatusResponse, error) {
	var err error
	const operationName = "LocationsClient.GetAzureAsyncOperationStatus"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getAzureAsyncOperationStatusCreateRequest(ctx, location, operationID, options)
	if err != nil {
		return LocationsClientGetAzureAsyncOperationStatusResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return LocationsClientGetAzureAsyncOperationStatusResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return LocationsClientGetAzureAsyncOperationStatusResponse{}, err
	}
	resp, err := client.getAzureAsyncOperationStatusHandleResponse(httpResp)
	return resp, err
}

// getAzureAsyncOperationStatusCreateRequest creates the GetAzureAsyncOperationStatus request.
func (client *LocationsClient) getAzureAsyncOperationStatusCreateRequest(ctx context.Context, location string, operationID string, options *LocationsClientGetAzureAsyncOperationStatusOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.HDInsight/locations/{location}/azureasyncoperations/{operationId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if operationID == "" {
		return nil, errors.New("parameter operationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{operationId}", url.PathEscape(operationID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-08-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getAzureAsyncOperationStatusHandleResponse handles the GetAzureAsyncOperationStatus response.
func (client *LocationsClient) getAzureAsyncOperationStatusHandleResponse(resp *http.Response) (LocationsClientGetAzureAsyncOperationStatusResponse, error) {
	result := LocationsClientGetAzureAsyncOperationStatusResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AsyncOperationResult); err != nil {
		return LocationsClientGetAzureAsyncOperationStatusResponse{}, err
	}
	return result, nil
}

// GetCapabilities - Gets the capabilities for the specified location.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-08-15-preview
//   - location - The Azure location (region) for which to make the request.
//   - options - LocationsClientGetCapabilitiesOptions contains the optional parameters for the LocationsClient.GetCapabilities
//     method.
func (client *LocationsClient) GetCapabilities(ctx context.Context, location string, options *LocationsClientGetCapabilitiesOptions) (LocationsClientGetCapabilitiesResponse, error) {
	var err error
	const operationName = "LocationsClient.GetCapabilities"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCapabilitiesCreateRequest(ctx, location, options)
	if err != nil {
		return LocationsClientGetCapabilitiesResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return LocationsClientGetCapabilitiesResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return LocationsClientGetCapabilitiesResponse{}, err
	}
	resp, err := client.getCapabilitiesHandleResponse(httpResp)
	return resp, err
}

// getCapabilitiesCreateRequest creates the GetCapabilities request.
func (client *LocationsClient) getCapabilitiesCreateRequest(ctx context.Context, location string, options *LocationsClientGetCapabilitiesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.HDInsight/locations/{location}/capabilities"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-08-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getCapabilitiesHandleResponse handles the GetCapabilities response.
func (client *LocationsClient) getCapabilitiesHandleResponse(resp *http.Response) (LocationsClientGetCapabilitiesResponse, error) {
	result := LocationsClientGetCapabilitiesResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CapabilitiesResult); err != nil {
		return LocationsClientGetCapabilitiesResponse{}, err
	}
	return result, nil
}

// ListBillingSpecs - Lists the billingSpecs for the specified subscription and location.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-08-15-preview
//   - location - The Azure location (region) for which to make the request.
//   - options - LocationsClientListBillingSpecsOptions contains the optional parameters for the LocationsClient.ListBillingSpecs
//     method.
func (client *LocationsClient) ListBillingSpecs(ctx context.Context, location string, options *LocationsClientListBillingSpecsOptions) (LocationsClientListBillingSpecsResponse, error) {
	var err error
	const operationName = "LocationsClient.ListBillingSpecs"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listBillingSpecsCreateRequest(ctx, location, options)
	if err != nil {
		return LocationsClientListBillingSpecsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return LocationsClientListBillingSpecsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return LocationsClientListBillingSpecsResponse{}, err
	}
	resp, err := client.listBillingSpecsHandleResponse(httpResp)
	return resp, err
}

// listBillingSpecsCreateRequest creates the ListBillingSpecs request.
func (client *LocationsClient) listBillingSpecsCreateRequest(ctx context.Context, location string, options *LocationsClientListBillingSpecsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.HDInsight/locations/{location}/billingSpecs"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-08-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBillingSpecsHandleResponse handles the ListBillingSpecs response.
func (client *LocationsClient) listBillingSpecsHandleResponse(resp *http.Response) (LocationsClientListBillingSpecsResponse, error) {
	result := LocationsClientListBillingSpecsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BillingResponseListResult); err != nil {
		return LocationsClientListBillingSpecsResponse{}, err
	}
	return result, nil
}

// ListUsages - Lists the usages for the specified location.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-08-15-preview
//   - location - The Azure location (region) for which to make the request.
//   - options - LocationsClientListUsagesOptions contains the optional parameters for the LocationsClient.ListUsages method.
func (client *LocationsClient) ListUsages(ctx context.Context, location string, options *LocationsClientListUsagesOptions) (LocationsClientListUsagesResponse, error) {
	var err error
	const operationName = "LocationsClient.ListUsages"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listUsagesCreateRequest(ctx, location, options)
	if err != nil {
		return LocationsClientListUsagesResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return LocationsClientListUsagesResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return LocationsClientListUsagesResponse{}, err
	}
	resp, err := client.listUsagesHandleResponse(httpResp)
	return resp, err
}

// listUsagesCreateRequest creates the ListUsages request.
func (client *LocationsClient) listUsagesCreateRequest(ctx context.Context, location string, options *LocationsClientListUsagesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.HDInsight/locations/{location}/usages"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-08-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listUsagesHandleResponse handles the ListUsages response.
func (client *LocationsClient) listUsagesHandleResponse(resp *http.Response) (LocationsClientListUsagesResponse, error) {
	result := LocationsClientListUsagesResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.UsagesListResult); err != nil {
		return LocationsClientListUsagesResponse{}, err
	}
	return result, nil
}

// ValidateClusterCreateRequest - Validate the cluster create request spec is valid or not.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-08-15-preview
//   - location - The Azure location (region) for which to make the request.
//   - options - LocationsClientValidateClusterCreateRequestOptions contains the optional parameters for the LocationsClient.ValidateClusterCreateRequest
//     method.
func (client *LocationsClient) ValidateClusterCreateRequest(ctx context.Context, location string, parameters ClusterCreateRequestValidationParameters, options *LocationsClientValidateClusterCreateRequestOptions) (LocationsClientValidateClusterCreateRequestResponse, error) {
	var err error
	const operationName = "LocationsClient.ValidateClusterCreateRequest"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.validateClusterCreateRequestCreateRequest(ctx, location, parameters, options)
	if err != nil {
		return LocationsClientValidateClusterCreateRequestResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return LocationsClientValidateClusterCreateRequestResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return LocationsClientValidateClusterCreateRequestResponse{}, err
	}
	resp, err := client.validateClusterCreateRequestHandleResponse(httpResp)
	return resp, err
}

// validateClusterCreateRequestCreateRequest creates the ValidateClusterCreateRequest request.
func (client *LocationsClient) validateClusterCreateRequestCreateRequest(ctx context.Context, location string, parameters ClusterCreateRequestValidationParameters, options *LocationsClientValidateClusterCreateRequestOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.HDInsight/locations/{location}/validateCreateRequest"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-08-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// validateClusterCreateRequestHandleResponse handles the ValidateClusterCreateRequest response.
func (client *LocationsClient) validateClusterCreateRequestHandleResponse(resp *http.Response) (LocationsClientValidateClusterCreateRequestResponse, error) {
	result := LocationsClientValidateClusterCreateRequestResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ClusterCreateValidationResult); err != nil {
		return LocationsClientValidateClusterCreateRequestResponse{}, err
	}
	return result, nil
}
