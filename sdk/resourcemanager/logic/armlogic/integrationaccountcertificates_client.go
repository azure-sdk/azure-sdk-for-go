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

// IntegrationAccountCertificatesClient contains the methods for the IntegrationAccountCertificates group.
// Don't use this type directly, use NewIntegrationAccountCertificatesClient() instead.
type IntegrationAccountCertificatesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewIntegrationAccountCertificatesClient creates a new instance of IntegrationAccountCertificatesClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewIntegrationAccountCertificatesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*IntegrationAccountCertificatesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &IntegrationAccountCertificatesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates an integration account certificate.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-05-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - integrationAccountName - The integration account name.
//   - certificateName - The integration account certificate name.
//   - resource - The integration account certificate.
//   - options - IntegrationAccountCertificatesClientCreateOrUpdateOptions contains the optional parameters for the IntegrationAccountCertificatesClient.CreateOrUpdate
//     method.
func (client *IntegrationAccountCertificatesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, integrationAccountName string, certificateName string, resource IntegrationAccountCertificate, options *IntegrationAccountCertificatesClientCreateOrUpdateOptions) (IntegrationAccountCertificatesClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "IntegrationAccountCertificatesClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, integrationAccountName, certificateName, resource, options)
	if err != nil {
		return IntegrationAccountCertificatesClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return IntegrationAccountCertificatesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return IntegrationAccountCertificatesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *IntegrationAccountCertificatesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, integrationAccountName string, certificateName string, resource IntegrationAccountCertificate, _ *IntegrationAccountCertificatesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/certificates/{certificateName}"
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
	if certificateName == "" {
		return nil, errors.New("parameter certificateName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{certificateName}", url.PathEscape(certificateName))
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
func (client *IntegrationAccountCertificatesClient) createOrUpdateHandleResponse(resp *http.Response) (IntegrationAccountCertificatesClientCreateOrUpdateResponse, error) {
	result := IntegrationAccountCertificatesClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IntegrationAccountCertificate); err != nil {
		return IntegrationAccountCertificatesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes an integration account certificate.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-05-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - integrationAccountName - The integration account name.
//   - certificateName - The integration account certificate name.
//   - options - IntegrationAccountCertificatesClientDeleteOptions contains the optional parameters for the IntegrationAccountCertificatesClient.Delete
//     method.
func (client *IntegrationAccountCertificatesClient) Delete(ctx context.Context, resourceGroupName string, integrationAccountName string, certificateName string, options *IntegrationAccountCertificatesClientDeleteOptions) (IntegrationAccountCertificatesClientDeleteResponse, error) {
	var err error
	const operationName = "IntegrationAccountCertificatesClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, integrationAccountName, certificateName, options)
	if err != nil {
		return IntegrationAccountCertificatesClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return IntegrationAccountCertificatesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return IntegrationAccountCertificatesClientDeleteResponse{}, err
	}
	return IntegrationAccountCertificatesClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *IntegrationAccountCertificatesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, integrationAccountName string, certificateName string, _ *IntegrationAccountCertificatesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/certificates/{certificateName}"
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
	if certificateName == "" {
		return nil, errors.New("parameter certificateName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{certificateName}", url.PathEscape(certificateName))
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

// Get - Gets an integration account certificate.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-05-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - integrationAccountName - The integration account name.
//   - certificateName - The integration account certificate name.
//   - options - IntegrationAccountCertificatesClientGetOptions contains the optional parameters for the IntegrationAccountCertificatesClient.Get
//     method.
func (client *IntegrationAccountCertificatesClient) Get(ctx context.Context, resourceGroupName string, integrationAccountName string, certificateName string, options *IntegrationAccountCertificatesClientGetOptions) (IntegrationAccountCertificatesClientGetResponse, error) {
	var err error
	const operationName = "IntegrationAccountCertificatesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, integrationAccountName, certificateName, options)
	if err != nil {
		return IntegrationAccountCertificatesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return IntegrationAccountCertificatesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return IntegrationAccountCertificatesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *IntegrationAccountCertificatesClient) getCreateRequest(ctx context.Context, resourceGroupName string, integrationAccountName string, certificateName string, _ *IntegrationAccountCertificatesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/certificates/{certificateName}"
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
	if certificateName == "" {
		return nil, errors.New("parameter certificateName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{certificateName}", url.PathEscape(certificateName))
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
func (client *IntegrationAccountCertificatesClient) getHandleResponse(resp *http.Response) (IntegrationAccountCertificatesClientGetResponse, error) {
	result := IntegrationAccountCertificatesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IntegrationAccountCertificate); err != nil {
		return IntegrationAccountCertificatesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets a list of integration account certificates.
//
// Generated from API version 2019-05-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - integrationAccountName - The integration account name.
//   - options - IntegrationAccountCertificatesClientListOptions contains the optional parameters for the IntegrationAccountCertificatesClient.NewListPager
//     method.
func (client *IntegrationAccountCertificatesClient) NewListPager(resourceGroupName string, integrationAccountName string, options *IntegrationAccountCertificatesClientListOptions) *runtime.Pager[IntegrationAccountCertificatesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[IntegrationAccountCertificatesClientListResponse]{
		More: func(page IntegrationAccountCertificatesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *IntegrationAccountCertificatesClientListResponse) (IntegrationAccountCertificatesClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "IntegrationAccountCertificatesClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, integrationAccountName, options)
			}, nil)
			if err != nil {
				return IntegrationAccountCertificatesClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *IntegrationAccountCertificatesClient) listCreateRequest(ctx context.Context, resourceGroupName string, integrationAccountName string, options *IntegrationAccountCertificatesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/certificates"
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
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *IntegrationAccountCertificatesClient) listHandleResponse(resp *http.Response) (IntegrationAccountCertificatesClientListResponse, error) {
	result := IntegrationAccountCertificatesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IntegrationAccountCertificateListResult); err != nil {
		return IntegrationAccountCertificatesClientListResponse{}, err
	}
	return result, nil
}
