//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armaddons

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

// SupportPlanTypesClient contains the methods for the SupportPlanTypes group.
// Don't use this type directly, use NewSupportPlanTypesClient() instead.
type SupportPlanTypesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewSupportPlanTypesClient creates a new instance of SupportPlanTypesClient with the specified values.
//   - subscriptionID - Subscription credentials that uniquely identify the Microsoft Azure subscription. The subscription ID
//     forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewSupportPlanTypesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SupportPlanTypesClient, error) {
	cl, err := arm.NewClient(moduleName+".SupportPlanTypesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &SupportPlanTypesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates the Canonical support plan of type {type} for the subscription.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-03-01
//   - providerName - The support plan type. For now the only valid type is "canonical".
//   - planTypeName - The Canonical support plan type.
//   - options - SupportPlanTypesClientBeginCreateOrUpdateOptions contains the optional parameters for the SupportPlanTypesClient.BeginCreateOrUpdate
//     method.
func (client *SupportPlanTypesClient) BeginCreateOrUpdate(ctx context.Context, providerName string, planTypeName PlanTypeName, options *SupportPlanTypesClientBeginCreateOrUpdateOptions) (*runtime.Poller[SupportPlanTypesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, providerName, planTypeName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[SupportPlanTypesClientCreateOrUpdateResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[SupportPlanTypesClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Creates or updates the Canonical support plan of type {type} for the subscription.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-03-01
func (client *SupportPlanTypesClient) createOrUpdate(ctx context.Context, providerName string, planTypeName PlanTypeName, options *SupportPlanTypesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	req, err := client.createOrUpdateCreateRequest(ctx, providerName, planTypeName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated, http.StatusNotFound) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *SupportPlanTypesClient) createOrUpdateCreateRequest(ctx context.Context, providerName string, planTypeName PlanTypeName, options *SupportPlanTypesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Addons/supportProviders/{providerName}/supportPlanTypes/{planTypeName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if providerName == "" {
		return nil, errors.New("parameter providerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{providerName}", url.PathEscape(providerName))
	if planTypeName == "" {
		return nil, errors.New("parameter planTypeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{planTypeName}", url.PathEscape(string(planTypeName)))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginDelete - Cancels the Canonical support plan of type {type} for the subscription.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-03-01
//   - providerName - The support plan type. For now the only valid type is "canonical".
//   - planTypeName - The Canonical support plan type.
//   - options - SupportPlanTypesClientBeginDeleteOptions contains the optional parameters for the SupportPlanTypesClient.BeginDelete
//     method.
func (client *SupportPlanTypesClient) BeginDelete(ctx context.Context, providerName string, planTypeName PlanTypeName, options *SupportPlanTypesClientBeginDeleteOptions) (*runtime.Poller[SupportPlanTypesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, providerName, planTypeName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[SupportPlanTypesClientDeleteResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[SupportPlanTypesClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Cancels the Canonical support plan of type {type} for the subscription.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-03-01
func (client *SupportPlanTypesClient) deleteOperation(ctx context.Context, providerName string, planTypeName PlanTypeName, options *SupportPlanTypesClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, providerName, planTypeName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *SupportPlanTypesClient) deleteCreateRequest(ctx context.Context, providerName string, planTypeName PlanTypeName, options *SupportPlanTypesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Addons/supportProviders/{providerName}/supportPlanTypes/{planTypeName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if providerName == "" {
		return nil, errors.New("parameter providerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{providerName}", url.PathEscape(providerName))
	if planTypeName == "" {
		return nil, errors.New("parameter planTypeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{planTypeName}", url.PathEscape(string(planTypeName)))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Returns whether or not the canonical support plan of type {type} is enabled for the subscription.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-03-01
//   - providerName - The support plan type. For now the only valid type is "canonical".
//   - planTypeName - The Canonical support plan type.
//   - options - SupportPlanTypesClientGetOptions contains the optional parameters for the SupportPlanTypesClient.Get method.
func (client *SupportPlanTypesClient) Get(ctx context.Context, providerName string, planTypeName PlanTypeName, options *SupportPlanTypesClientGetOptions) (SupportPlanTypesClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, providerName, planTypeName, options)
	if err != nil {
		return SupportPlanTypesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SupportPlanTypesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNotFound) {
		err = runtime.NewResponseError(httpResp)
		return SupportPlanTypesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *SupportPlanTypesClient) getCreateRequest(ctx context.Context, providerName string, planTypeName PlanTypeName, options *SupportPlanTypesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Addons/supportProviders/{providerName}/supportPlanTypes/{planTypeName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if providerName == "" {
		return nil, errors.New("parameter providerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{providerName}", url.PathEscape(providerName))
	if planTypeName == "" {
		return nil, errors.New("parameter planTypeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{planTypeName}", url.PathEscape(string(planTypeName)))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SupportPlanTypesClient) getHandleResponse(resp *http.Response) (SupportPlanTypesClientGetResponse, error) {
	result := SupportPlanTypesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CanonicalSupportPlanResponseEnvelope); err != nil {
		return SupportPlanTypesClientGetResponse{}, err
	}
	return result, nil
}

// ListInfo - Returns the canonical support plan information for all types for the subscription.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-03-01
//   - options - SupportPlanTypesClientListInfoOptions contains the optional parameters for the SupportPlanTypesClient.ListInfo
//     method.
func (client *SupportPlanTypesClient) ListInfo(ctx context.Context, options *SupportPlanTypesClientListInfoOptions) (SupportPlanTypesClientListInfoResponse, error) {
	var err error
	req, err := client.listInfoCreateRequest(ctx, options)
	if err != nil {
		return SupportPlanTypesClientListInfoResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SupportPlanTypesClientListInfoResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNotFound) {
		err = runtime.NewResponseError(httpResp)
		return SupportPlanTypesClientListInfoResponse{}, err
	}
	resp, err := client.listInfoHandleResponse(httpResp)
	return resp, err
}

// listInfoCreateRequest creates the ListInfo request.
func (client *SupportPlanTypesClient) listInfoCreateRequest(ctx context.Context, options *SupportPlanTypesClientListInfoOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Addons/supportProviders/canonical/listSupportPlanInfo"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listInfoHandleResponse handles the ListInfo response.
func (client *SupportPlanTypesClient) listInfoHandleResponse(resp *http.Response) (SupportPlanTypesClientListInfoResponse, error) {
	result := SupportPlanTypesClientListInfoResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CanonicalSupportPlanInfoDefinitionArray); err != nil {
		return SupportPlanTypesClientListInfoResponse{}, err
	}
	return result, nil
}
