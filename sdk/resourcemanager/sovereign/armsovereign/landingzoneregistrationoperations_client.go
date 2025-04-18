// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsovereign

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

// LandingZoneRegistrationOperationsClient contains the methods for the LandingZoneRegistrationOperations group.
// Don't use this type directly, use NewLandingZoneRegistrationOperationsClient() instead.
type LandingZoneRegistrationOperationsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewLandingZoneRegistrationOperationsClient creates a new instance of LandingZoneRegistrationOperationsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewLandingZoneRegistrationOperationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*LandingZoneRegistrationOperationsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &LandingZoneRegistrationOperationsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreate - Create a landing zone registration.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-02-27-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - landingZoneAccountName - The landing zone account.
//   - landingZoneRegistrationName - The name of the landing zone registration resource.
//   - resource - Resource create parameters.
//   - options - LandingZoneRegistrationOperationsClientBeginCreateOptions contains the optional parameters for the LandingZoneRegistrationOperationsClient.BeginCreate
//     method.
func (client *LandingZoneRegistrationOperationsClient) BeginCreate(ctx context.Context, resourceGroupName string, landingZoneAccountName string, landingZoneRegistrationName string, resource LandingZoneRegistrationResource, options *LandingZoneRegistrationOperationsClientBeginCreateOptions) (*runtime.Poller[LandingZoneRegistrationOperationsClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, landingZoneAccountName, landingZoneRegistrationName, resource, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[LandingZoneRegistrationOperationsClientCreateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[LandingZoneRegistrationOperationsClientCreateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Create - Create a landing zone registration.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-02-27-preview
func (client *LandingZoneRegistrationOperationsClient) create(ctx context.Context, resourceGroupName string, landingZoneAccountName string, landingZoneRegistrationName string, resource LandingZoneRegistrationResource, options *LandingZoneRegistrationOperationsClientBeginCreateOptions) (*http.Response, error) {
	var err error
	const operationName = "LandingZoneRegistrationOperationsClient.BeginCreate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createCreateRequest(ctx, resourceGroupName, landingZoneAccountName, landingZoneRegistrationName, resource, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createCreateRequest creates the Create request.
func (client *LandingZoneRegistrationOperationsClient) createCreateRequest(ctx context.Context, resourceGroupName string, landingZoneAccountName string, landingZoneRegistrationName string, resource LandingZoneRegistrationResource, _ *LandingZoneRegistrationOperationsClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sovereign/landingZoneAccounts/{landingZoneAccountName}/landingZoneRegistrations/{landingZoneRegistrationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if landingZoneAccountName == "" {
		return nil, errors.New("parameter landingZoneAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{landingZoneAccountName}", url.PathEscape(landingZoneAccountName))
	if landingZoneRegistrationName == "" {
		return nil, errors.New("parameter landingZoneRegistrationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{landingZoneRegistrationName}", url.PathEscape(landingZoneRegistrationName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-02-27-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, resource); err != nil {
		return nil, err
	}
	return req, nil
}

// Delete - Delete a landing zone registration.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-02-27-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - landingZoneAccountName - The landing zone account.
//   - landingZoneRegistrationName - The name of the landing zone registration resource.
//   - options - LandingZoneRegistrationOperationsClientDeleteOptions contains the optional parameters for the LandingZoneRegistrationOperationsClient.Delete
//     method.
func (client *LandingZoneRegistrationOperationsClient) Delete(ctx context.Context, resourceGroupName string, landingZoneAccountName string, landingZoneRegistrationName string, options *LandingZoneRegistrationOperationsClientDeleteOptions) (LandingZoneRegistrationOperationsClientDeleteResponse, error) {
	var err error
	const operationName = "LandingZoneRegistrationOperationsClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, landingZoneAccountName, landingZoneRegistrationName, options)
	if err != nil {
		return LandingZoneRegistrationOperationsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return LandingZoneRegistrationOperationsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return LandingZoneRegistrationOperationsClientDeleteResponse{}, err
	}
	return LandingZoneRegistrationOperationsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *LandingZoneRegistrationOperationsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, landingZoneAccountName string, landingZoneRegistrationName string, _ *LandingZoneRegistrationOperationsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sovereign/landingZoneAccounts/{landingZoneAccountName}/landingZoneRegistrations/{landingZoneRegistrationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if landingZoneAccountName == "" {
		return nil, errors.New("parameter landingZoneAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{landingZoneAccountName}", url.PathEscape(landingZoneAccountName))
	if landingZoneRegistrationName == "" {
		return nil, errors.New("parameter landingZoneRegistrationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{landingZoneRegistrationName}", url.PathEscape(landingZoneRegistrationName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-02-27-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get a landing zone registration.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-02-27-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - landingZoneAccountName - The landing zone account.
//   - landingZoneRegistrationName - The name of the landing zone registration resource.
//   - options - LandingZoneRegistrationOperationsClientGetOptions contains the optional parameters for the LandingZoneRegistrationOperationsClient.Get
//     method.
func (client *LandingZoneRegistrationOperationsClient) Get(ctx context.Context, resourceGroupName string, landingZoneAccountName string, landingZoneRegistrationName string, options *LandingZoneRegistrationOperationsClientGetOptions) (LandingZoneRegistrationOperationsClientGetResponse, error) {
	var err error
	const operationName = "LandingZoneRegistrationOperationsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, landingZoneAccountName, landingZoneRegistrationName, options)
	if err != nil {
		return LandingZoneRegistrationOperationsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return LandingZoneRegistrationOperationsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return LandingZoneRegistrationOperationsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *LandingZoneRegistrationOperationsClient) getCreateRequest(ctx context.Context, resourceGroupName string, landingZoneAccountName string, landingZoneRegistrationName string, _ *LandingZoneRegistrationOperationsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sovereign/landingZoneAccounts/{landingZoneAccountName}/landingZoneRegistrations/{landingZoneRegistrationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if landingZoneAccountName == "" {
		return nil, errors.New("parameter landingZoneAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{landingZoneAccountName}", url.PathEscape(landingZoneAccountName))
	if landingZoneRegistrationName == "" {
		return nil, errors.New("parameter landingZoneRegistrationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{landingZoneRegistrationName}", url.PathEscape(landingZoneRegistrationName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-02-27-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *LandingZoneRegistrationOperationsClient) getHandleResponse(resp *http.Response) (LandingZoneRegistrationOperationsClientGetResponse, error) {
	result := LandingZoneRegistrationOperationsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LandingZoneRegistrationResource); err != nil {
		return LandingZoneRegistrationOperationsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List the landing zone registrations within a landing zone account.
//
// Generated from API version 2025-02-27-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - landingZoneAccountName - The landing zone account.
//   - options - LandingZoneRegistrationOperationsClientListOptions contains the optional parameters for the LandingZoneRegistrationOperationsClient.NewListPager
//     method.
func (client *LandingZoneRegistrationOperationsClient) NewListPager(resourceGroupName string, landingZoneAccountName string, options *LandingZoneRegistrationOperationsClientListOptions) *runtime.Pager[LandingZoneRegistrationOperationsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[LandingZoneRegistrationOperationsClientListResponse]{
		More: func(page LandingZoneRegistrationOperationsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *LandingZoneRegistrationOperationsClientListResponse) (LandingZoneRegistrationOperationsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "LandingZoneRegistrationOperationsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, landingZoneAccountName, options)
			}, nil)
			if err != nil {
				return LandingZoneRegistrationOperationsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *LandingZoneRegistrationOperationsClient) listCreateRequest(ctx context.Context, resourceGroupName string, landingZoneAccountName string, _ *LandingZoneRegistrationOperationsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sovereign/landingZoneAccounts/{landingZoneAccountName}/landingZoneRegistrations"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if landingZoneAccountName == "" {
		return nil, errors.New("parameter landingZoneAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{landingZoneAccountName}", url.PathEscape(landingZoneAccountName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-02-27-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *LandingZoneRegistrationOperationsClient) listHandleResponse(resp *http.Response) (LandingZoneRegistrationOperationsClientListResponse, error) {
	result := LandingZoneRegistrationOperationsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LandingZoneRegistrationResourceListResult); err != nil {
		return LandingZoneRegistrationOperationsClientListResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Update a landing zone registration.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-02-27-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - landingZoneAccountName - The landing zone account.
//   - landingZoneRegistrationName - The name of the landing zone registration resource.
//   - properties - The resource properties to be updated.
//   - options - LandingZoneRegistrationOperationsClientBeginUpdateOptions contains the optional parameters for the LandingZoneRegistrationOperationsClient.BeginUpdate
//     method.
func (client *LandingZoneRegistrationOperationsClient) BeginUpdate(ctx context.Context, resourceGroupName string, landingZoneAccountName string, landingZoneRegistrationName string, properties LandingZoneRegistrationResourceUpdate, options *LandingZoneRegistrationOperationsClientBeginUpdateOptions) (*runtime.Poller[LandingZoneRegistrationOperationsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, landingZoneAccountName, landingZoneRegistrationName, properties, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[LandingZoneRegistrationOperationsClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[LandingZoneRegistrationOperationsClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Update - Update a landing zone registration.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-02-27-preview
func (client *LandingZoneRegistrationOperationsClient) update(ctx context.Context, resourceGroupName string, landingZoneAccountName string, landingZoneRegistrationName string, properties LandingZoneRegistrationResourceUpdate, options *LandingZoneRegistrationOperationsClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "LandingZoneRegistrationOperationsClient.BeginUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, landingZoneAccountName, landingZoneRegistrationName, properties, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// updateCreateRequest creates the Update request.
func (client *LandingZoneRegistrationOperationsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, landingZoneAccountName string, landingZoneRegistrationName string, properties LandingZoneRegistrationResourceUpdate, _ *LandingZoneRegistrationOperationsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sovereign/landingZoneAccounts/{landingZoneAccountName}/landingZoneRegistrations/{landingZoneRegistrationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if landingZoneAccountName == "" {
		return nil, errors.New("parameter landingZoneAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{landingZoneAccountName}", url.PathEscape(landingZoneAccountName))
	if landingZoneRegistrationName == "" {
		return nil, errors.New("parameter landingZoneRegistrationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{landingZoneRegistrationName}", url.PathEscape(landingZoneRegistrationName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-02-27-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, properties); err != nil {
		return nil, err
	}
	return req, nil
}
