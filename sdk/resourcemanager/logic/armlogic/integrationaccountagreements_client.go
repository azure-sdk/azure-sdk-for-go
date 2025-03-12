// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armlogic

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// IntegrationAccountAgreementsClient contains the methods for the IntegrationAccountAgreements group.
// Don't use this type directly, use NewIntegrationAccountAgreementsClient() instead.
type IntegrationAccountAgreementsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewIntegrationAccountAgreementsClient creates a new instance of IntegrationAccountAgreementsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewIntegrationAccountAgreementsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*IntegrationAccountAgreementsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &IntegrationAccountAgreementsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates an integration account agreement.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-05-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - integrationAccountName - The integration account name.
//   - agreementName - The integration account agreement name.
//   - resource - The integration account agreement.
//   - options - IntegrationAccountAgreementsClientCreateOrUpdateOptions contains the optional parameters for the IntegrationAccountAgreementsClient.CreateOrUpdate
//     method.
func (client *IntegrationAccountAgreementsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, integrationAccountName string, agreementName string, resource IntegrationAccountAgreement, options *IntegrationAccountAgreementsClientCreateOrUpdateOptions) (IntegrationAccountAgreementsClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "IntegrationAccountAgreementsClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, integrationAccountName, agreementName, resource, options)
	if err != nil {
		return IntegrationAccountAgreementsClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return IntegrationAccountAgreementsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return IntegrationAccountAgreementsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *IntegrationAccountAgreementsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, integrationAccountName string, agreementName string, resource IntegrationAccountAgreement, _ *IntegrationAccountAgreementsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/agreements/{agreementName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if integrationAccountName == "" {
		return nil, errors.New("parameter integrationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationAccountName}", url.PathEscape(integrationAccountName))
	if agreementName == "" {
		return nil, errors.New("parameter agreementName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{agreementName}", url.PathEscape(agreementName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, resource); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *IntegrationAccountAgreementsClient) createOrUpdateHandleResponse(resp *http.Response) (IntegrationAccountAgreementsClientCreateOrUpdateResponse, error) {
	result := IntegrationAccountAgreementsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IntegrationAccountAgreement); err != nil {
		return IntegrationAccountAgreementsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes an integration account agreement.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-05-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - integrationAccountName - The integration account name.
//   - agreementName - The integration account agreement name.
//   - options - IntegrationAccountAgreementsClientDeleteOptions contains the optional parameters for the IntegrationAccountAgreementsClient.Delete
//     method.
func (client *IntegrationAccountAgreementsClient) Delete(ctx context.Context, resourceGroupName string, integrationAccountName string, agreementName string, options *IntegrationAccountAgreementsClientDeleteOptions) (IntegrationAccountAgreementsClientDeleteResponse, error) {
	var err error
	const operationName = "IntegrationAccountAgreementsClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, integrationAccountName, agreementName, options)
	if err != nil {
		return IntegrationAccountAgreementsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return IntegrationAccountAgreementsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return IntegrationAccountAgreementsClientDeleteResponse{}, err
	}
	return IntegrationAccountAgreementsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *IntegrationAccountAgreementsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, integrationAccountName string, agreementName string, _ *IntegrationAccountAgreementsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/agreements/{agreementName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if integrationAccountName == "" {
		return nil, errors.New("parameter integrationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationAccountName}", url.PathEscape(integrationAccountName))
	if agreementName == "" {
		return nil, errors.New("parameter agreementName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{agreementName}", url.PathEscape(agreementName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets an integration account agreement.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-05-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - integrationAccountName - The integration account name.
//   - agreementName - The integration account agreement name.
//   - options - IntegrationAccountAgreementsClientGetOptions contains the optional parameters for the IntegrationAccountAgreementsClient.Get
//     method.
func (client *IntegrationAccountAgreementsClient) Get(ctx context.Context, resourceGroupName string, integrationAccountName string, agreementName string, options *IntegrationAccountAgreementsClientGetOptions) (IntegrationAccountAgreementsClientGetResponse, error) {
	var err error
	const operationName = "IntegrationAccountAgreementsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, integrationAccountName, agreementName, options)
	if err != nil {
		return IntegrationAccountAgreementsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return IntegrationAccountAgreementsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return IntegrationAccountAgreementsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *IntegrationAccountAgreementsClient) getCreateRequest(ctx context.Context, resourceGroupName string, integrationAccountName string, agreementName string, _ *IntegrationAccountAgreementsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/agreements/{agreementName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if integrationAccountName == "" {
		return nil, errors.New("parameter integrationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationAccountName}", url.PathEscape(integrationAccountName))
	if agreementName == "" {
		return nil, errors.New("parameter agreementName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{agreementName}", url.PathEscape(agreementName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *IntegrationAccountAgreementsClient) getHandleResponse(resp *http.Response) (IntegrationAccountAgreementsClientGetResponse, error) {
	result := IntegrationAccountAgreementsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IntegrationAccountAgreement); err != nil {
		return IntegrationAccountAgreementsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets a list of integration account agreements.
//
// Generated from API version 2019-05-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - integrationAccountName - The integration account name.
//   - options - IntegrationAccountAgreementsClientListOptions contains the optional parameters for the IntegrationAccountAgreementsClient.NewListPager
//     method.
func (client *IntegrationAccountAgreementsClient) NewListPager(resourceGroupName string, integrationAccountName string, options *IntegrationAccountAgreementsClientListOptions) *runtime.Pager[IntegrationAccountAgreementsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[IntegrationAccountAgreementsClientListResponse]{
		More: func(page IntegrationAccountAgreementsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *IntegrationAccountAgreementsClientListResponse) (IntegrationAccountAgreementsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "IntegrationAccountAgreementsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, integrationAccountName, options)
			}, nil)
			if err != nil {
				return IntegrationAccountAgreementsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *IntegrationAccountAgreementsClient) listCreateRequest(ctx context.Context, resourceGroupName string, integrationAccountName string, options *IntegrationAccountAgreementsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/agreements"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if integrationAccountName == "" {
		return nil, errors.New("parameter integrationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationAccountName}", url.PathEscape(integrationAccountName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *IntegrationAccountAgreementsClient) listHandleResponse(resp *http.Response) (IntegrationAccountAgreementsClientListResponse, error) {
	result := IntegrationAccountAgreementsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IntegrationAccountAgreementListResult); err != nil {
		return IntegrationAccountAgreementsClientListResponse{}, err
	}
	return result, nil
}

// ListContentCallbackURL - Get the content callback url.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-05-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - integrationAccountName - The integration account name.
//   - agreementName - The integration account agreement name.
//   - options - IntegrationAccountAgreementsClientListContentCallbackURLOptions contains the optional parameters for the IntegrationAccountAgreementsClient.ListContentCallbackURL
//     method.
func (client *IntegrationAccountAgreementsClient) ListContentCallbackURL(ctx context.Context, resourceGroupName string, integrationAccountName string, agreementName string, body GetCallbackURLParameters, options *IntegrationAccountAgreementsClientListContentCallbackURLOptions) (IntegrationAccountAgreementsClientListContentCallbackURLResponse, error) {
	var err error
	const operationName = "IntegrationAccountAgreementsClient.ListContentCallbackURL"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listContentCallbackURLCreateRequest(ctx, resourceGroupName, integrationAccountName, agreementName, body, options)
	if err != nil {
		return IntegrationAccountAgreementsClientListContentCallbackURLResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return IntegrationAccountAgreementsClientListContentCallbackURLResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return IntegrationAccountAgreementsClientListContentCallbackURLResponse{}, err
	}
	resp, err := client.listContentCallbackURLHandleResponse(httpResp)
	return resp, err
}

// listContentCallbackURLCreateRequest creates the ListContentCallbackURL request.
func (client *IntegrationAccountAgreementsClient) listContentCallbackURLCreateRequest(ctx context.Context, resourceGroupName string, integrationAccountName string, agreementName string, body GetCallbackURLParameters, _ *IntegrationAccountAgreementsClientListContentCallbackURLOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/agreements/{agreementName}/listContentCallbackUrl"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if integrationAccountName == "" {
		return nil, errors.New("parameter integrationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationAccountName}", url.PathEscape(integrationAccountName))
	if agreementName == "" {
		return nil, errors.New("parameter agreementName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{agreementName}", url.PathEscape(agreementName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// listContentCallbackURLHandleResponse handles the ListContentCallbackURL response.
func (client *IntegrationAccountAgreementsClient) listContentCallbackURLHandleResponse(resp *http.Response) (IntegrationAccountAgreementsClientListContentCallbackURLResponse, error) {
	result := IntegrationAccountAgreementsClientListContentCallbackURLResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WorkflowTriggerCallbackURL); err != nil {
		return IntegrationAccountAgreementsClientListContentCallbackURLResponse{}, err
	}
	return result, nil
}
