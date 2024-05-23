//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsecurity

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

// TrustedIPsClient contains the methods for the TrustedIPs group.
// Don't use this type directly, use NewTrustedIPsClient() instead.
type TrustedIPsClient struct {
	internal *arm.Client
}

// NewTrustedIPsClient creates a new instance of TrustedIPsClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewTrustedIPsClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*TrustedIPsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &TrustedIPsClient{
		internal: cl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates a trusted IPs object over a given scope
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-07-01-preview
//   - scope - The scope of the Trusted IPs. Valid scopes are: management group (format: 'providers/Microsoft.Management/managementGroups/{managementGroup}'),
//     subscription (format: 'subscriptions/{subscriptionId}'),
//     or security connector (format: 'subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/securityConnectors/{securityConnectorName})'
//   - trustedIPsID - The trusted IPs key - unique key for the specific trusted IPs range (GUID)
//   - trustedIPs - Trusted IPs over a given scope
//   - options - TrustedIPsClientCreateOrUpdateOptions contains the optional parameters for the TrustedIPsClient.CreateOrUpdate
//     method.
func (client *TrustedIPsClient) CreateOrUpdate(ctx context.Context, scope string, trustedIPsID string, trustedIPs TrustedIPs, options *TrustedIPsClientCreateOrUpdateOptions) (TrustedIPsClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "TrustedIPsClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, scope, trustedIPsID, trustedIPs, options)
	if err != nil {
		return TrustedIPsClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TrustedIPsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return TrustedIPsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *TrustedIPsClient) createOrUpdateCreateRequest(ctx context.Context, scope string, trustedIPsID string, trustedIPs TrustedIPs, options *TrustedIPsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Security/trustedIps/{trustedIpsId}"
	if scope == "" {
		return nil, errors.New("parameter scope cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scope}", url.PathEscape(scope))
	if trustedIPsID == "" {
		return nil, errors.New("parameter trustedIPsID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{trustedIpsId}", url.PathEscape(trustedIPsID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, trustedIPs); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *TrustedIPsClient) createOrUpdateHandleResponse(resp *http.Response) (TrustedIPsClientCreateOrUpdateResponse, error) {
	result := TrustedIPsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TrustedIPs); err != nil {
		return TrustedIPsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete a trusted IPs object over a given scope
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-07-01-preview
//   - scope - The scope of the Trusted IPs. Valid scopes are: management group (format: 'providers/Microsoft.Management/managementGroups/{managementGroup}'),
//     subscription (format: 'subscriptions/{subscriptionId}'),
//     or security connector (format: 'subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/securityConnectors/{securityConnectorName})'
//   - trustedIPsID - The trusted IPs key - unique key for the specific trusted IPs range (GUID)
//   - options - TrustedIPsClientDeleteOptions contains the optional parameters for the TrustedIPsClient.Delete method.
func (client *TrustedIPsClient) Delete(ctx context.Context, scope string, trustedIPsID string, options *TrustedIPsClientDeleteOptions) (TrustedIPsClientDeleteResponse, error) {
	var err error
	const operationName = "TrustedIPsClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, scope, trustedIPsID, options)
	if err != nil {
		return TrustedIPsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TrustedIPsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return TrustedIPsClientDeleteResponse{}, err
	}
	return TrustedIPsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *TrustedIPsClient) deleteCreateRequest(ctx context.Context, scope string, trustedIPsID string, options *TrustedIPsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Security/trustedIps/{trustedIpsId}"
	if scope == "" {
		return nil, errors.New("parameter scope cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scope}", url.PathEscape(scope))
	if trustedIPsID == "" {
		return nil, errors.New("parameter trustedIPsID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{trustedIpsId}", url.PathEscape(trustedIPsID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get a specific trusted IPs object for the requested scope by trusted IPs Id
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-07-01-preview
//   - scope - The scope of the Trusted IPs. Valid scopes are: management group (format: 'providers/Microsoft.Management/managementGroups/{managementGroup}'),
//     subscription (format: 'subscriptions/{subscriptionId}'),
//     or security connector (format: 'subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/securityConnectors/{securityConnectorName})'
//   - trustedIPsID - The trusted IPs key - unique key for the specific trusted IPs range (GUID)
//   - options - TrustedIPsClientGetOptions contains the optional parameters for the TrustedIPsClient.Get method.
func (client *TrustedIPsClient) Get(ctx context.Context, scope string, trustedIPsID string, options *TrustedIPsClientGetOptions) (TrustedIPsClientGetResponse, error) {
	var err error
	const operationName = "TrustedIPsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, scope, trustedIPsID, options)
	if err != nil {
		return TrustedIPsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TrustedIPsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return TrustedIPsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *TrustedIPsClient) getCreateRequest(ctx context.Context, scope string, trustedIPsID string, options *TrustedIPsClientGetOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Security/trustedIps/{trustedIpsId}"
	if scope == "" {
		return nil, errors.New("parameter scope cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scope}", url.PathEscape(scope))
	if trustedIPsID == "" {
		return nil, errors.New("parameter trustedIPsID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{trustedIpsId}", url.PathEscape(trustedIPsID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *TrustedIPsClient) getHandleResponse(resp *http.Response) (TrustedIPsClientGetResponse, error) {
	result := TrustedIPsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TrustedIPs); err != nil {
		return TrustedIPsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Get a list of all relevant trusted IPs over a scope
//
// Generated from API version 2024-07-01-preview
//   - scope - The scope of the Trusted IPs. Valid scopes are: management group (format: 'providers/Microsoft.Management/managementGroups/{managementGroup}'),
//     subscription (format: 'subscriptions/{subscriptionId}'),
//     or security connector (format: 'subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/securityConnectors/{securityConnectorName})'
//   - options - TrustedIPsClientListOptions contains the optional parameters for the TrustedIPsClient.NewListPager method.
func (client *TrustedIPsClient) NewListPager(scope string, options *TrustedIPsClientListOptions) *runtime.Pager[TrustedIPsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[TrustedIPsClientListResponse]{
		More: func(page TrustedIPsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *TrustedIPsClientListResponse) (TrustedIPsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "TrustedIPsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, scope, options)
			}, nil)
			if err != nil {
				return TrustedIPsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *TrustedIPsClient) listCreateRequest(ctx context.Context, scope string, options *TrustedIPsClientListOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Security/trustedIps"
	if scope == "" {
		return nil, errors.New("parameter scope cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scope}", url.PathEscape(scope))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *TrustedIPsClient) listHandleResponse(resp *http.Response) (TrustedIPsClientListResponse, error) {
	result := TrustedIPsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TrustedIPsList); err != nil {
		return TrustedIPsClientListResponse{}, err
	}
	return result, nil
}
