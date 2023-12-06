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

// AzureDevOpsOrgsClient contains the methods for the AzureDevOpsOrgs group.
// Don't use this type directly, use NewAzureDevOpsOrgsClient() instead.
type AzureDevOpsOrgsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewAzureDevOpsOrgsClient creates a new instance of AzureDevOpsOrgsClient with the specified values.
//   - subscriptionID - Azure subscription ID
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewAzureDevOpsOrgsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AzureDevOpsOrgsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &AzureDevOpsOrgsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates monitored Azure DevOps organization details.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - securityConnectorName - The security connector name.
//   - orgName - The Azure DevOps organization name.
//   - azureDevOpsOrg - The Azure DevOps organization resource payload.
//   - options - AzureDevOpsOrgsClientBeginCreateOrUpdateOptions contains the optional parameters for the AzureDevOpsOrgsClient.BeginCreateOrUpdate
//     method.
func (client *AzureDevOpsOrgsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, securityConnectorName string, orgName string, azureDevOpsOrg AzureDevOpsOrg, options *AzureDevOpsOrgsClientBeginCreateOrUpdateOptions) (*runtime.Poller[AzureDevOpsOrgsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, securityConnectorName, orgName, azureDevOpsOrg, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[AzureDevOpsOrgsClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[AzureDevOpsOrgsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Creates or updates monitored Azure DevOps organization details.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01-preview
func (client *AzureDevOpsOrgsClient) createOrUpdate(ctx context.Context, resourceGroupName string, securityConnectorName string, orgName string, azureDevOpsOrg AzureDevOpsOrg, options *AzureDevOpsOrgsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "AzureDevOpsOrgsClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, securityConnectorName, orgName, azureDevOpsOrg, options)
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

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *AzureDevOpsOrgsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, securityConnectorName string, orgName string, azureDevOpsOrg AzureDevOpsOrg, options *AzureDevOpsOrgsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/securityConnectors/{securityConnectorName}/devops/default/azureDevOpsOrgs/{orgName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if securityConnectorName == "" {
		return nil, errors.New("parameter securityConnectorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{securityConnectorName}", url.PathEscape(securityConnectorName))
	if orgName == "" {
		return nil, errors.New("parameter orgName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{orgName}", url.PathEscape(orgName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, azureDevOpsOrg); err != nil {
		return nil, err
	}
	return req, nil
}

// Get - Returns a monitored Azure DevOps organization resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - securityConnectorName - The security connector name.
//   - orgName - The Azure DevOps organization name.
//   - options - AzureDevOpsOrgsClientGetOptions contains the optional parameters for the AzureDevOpsOrgsClient.Get method.
func (client *AzureDevOpsOrgsClient) Get(ctx context.Context, resourceGroupName string, securityConnectorName string, orgName string, options *AzureDevOpsOrgsClientGetOptions) (AzureDevOpsOrgsClientGetResponse, error) {
	var err error
	const operationName = "AzureDevOpsOrgsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, securityConnectorName, orgName, options)
	if err != nil {
		return AzureDevOpsOrgsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AzureDevOpsOrgsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return AzureDevOpsOrgsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *AzureDevOpsOrgsClient) getCreateRequest(ctx context.Context, resourceGroupName string, securityConnectorName string, orgName string, options *AzureDevOpsOrgsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/securityConnectors/{securityConnectorName}/devops/default/azureDevOpsOrgs/{orgName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if securityConnectorName == "" {
		return nil, errors.New("parameter securityConnectorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{securityConnectorName}", url.PathEscape(securityConnectorName))
	if orgName == "" {
		return nil, errors.New("parameter orgName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{orgName}", url.PathEscape(orgName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AzureDevOpsOrgsClient) getHandleResponse(resp *http.Response) (AzureDevOpsOrgsClientGetResponse, error) {
	result := AzureDevOpsOrgsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AzureDevOpsOrg); err != nil {
		return AzureDevOpsOrgsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Returns a list of Azure DevOps organizations onboarded to the connector.
//
// Generated from API version 2023-09-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - securityConnectorName - The security connector name.
//   - options - AzureDevOpsOrgsClientListOptions contains the optional parameters for the AzureDevOpsOrgsClient.NewListPager
//     method.
func (client *AzureDevOpsOrgsClient) NewListPager(resourceGroupName string, securityConnectorName string, options *AzureDevOpsOrgsClientListOptions) *runtime.Pager[AzureDevOpsOrgsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[AzureDevOpsOrgsClientListResponse]{
		More: func(page AzureDevOpsOrgsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AzureDevOpsOrgsClientListResponse) (AzureDevOpsOrgsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "AzureDevOpsOrgsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, securityConnectorName, options)
			}, nil)
			if err != nil {
				return AzureDevOpsOrgsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *AzureDevOpsOrgsClient) listCreateRequest(ctx context.Context, resourceGroupName string, securityConnectorName string, options *AzureDevOpsOrgsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/securityConnectors/{securityConnectorName}/devops/default/azureDevOpsOrgs"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if securityConnectorName == "" {
		return nil, errors.New("parameter securityConnectorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{securityConnectorName}", url.PathEscape(securityConnectorName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *AzureDevOpsOrgsClient) listHandleResponse(resp *http.Response) (AzureDevOpsOrgsClientListResponse, error) {
	result := AzureDevOpsOrgsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AzureDevOpsOrgListResponse); err != nil {
		return AzureDevOpsOrgsClientListResponse{}, err
	}
	return result, nil
}

// ListAvailable - Returns a list of all Azure DevOps organizations accessible by the user token consumed by the connector.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - securityConnectorName - The security connector name.
//   - options - AzureDevOpsOrgsClientListAvailableOptions contains the optional parameters for the AzureDevOpsOrgsClient.ListAvailable
//     method.
func (client *AzureDevOpsOrgsClient) ListAvailable(ctx context.Context, resourceGroupName string, securityConnectorName string, options *AzureDevOpsOrgsClientListAvailableOptions) (AzureDevOpsOrgsClientListAvailableResponse, error) {
	var err error
	const operationName = "AzureDevOpsOrgsClient.ListAvailable"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listAvailableCreateRequest(ctx, resourceGroupName, securityConnectorName, options)
	if err != nil {
		return AzureDevOpsOrgsClientListAvailableResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AzureDevOpsOrgsClientListAvailableResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return AzureDevOpsOrgsClientListAvailableResponse{}, err
	}
	resp, err := client.listAvailableHandleResponse(httpResp)
	return resp, err
}

// listAvailableCreateRequest creates the ListAvailable request.
func (client *AzureDevOpsOrgsClient) listAvailableCreateRequest(ctx context.Context, resourceGroupName string, securityConnectorName string, options *AzureDevOpsOrgsClientListAvailableOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/securityConnectors/{securityConnectorName}/devops/default/listAvailableAzureDevOpsOrgs"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if securityConnectorName == "" {
		return nil, errors.New("parameter securityConnectorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{securityConnectorName}", url.PathEscape(securityConnectorName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listAvailableHandleResponse handles the ListAvailable response.
func (client *AzureDevOpsOrgsClient) listAvailableHandleResponse(resp *http.Response) (AzureDevOpsOrgsClientListAvailableResponse, error) {
	result := AzureDevOpsOrgsClientListAvailableResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AzureDevOpsOrgListResponse); err != nil {
		return AzureDevOpsOrgsClientListAvailableResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Updates monitored Azure DevOps organization details.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - securityConnectorName - The security connector name.
//   - orgName - The Azure DevOps organization name.
//   - azureDevOpsOrg - The Azure DevOps organization resource payload.
//   - options - AzureDevOpsOrgsClientBeginUpdateOptions contains the optional parameters for the AzureDevOpsOrgsClient.BeginUpdate
//     method.
func (client *AzureDevOpsOrgsClient) BeginUpdate(ctx context.Context, resourceGroupName string, securityConnectorName string, orgName string, azureDevOpsOrg AzureDevOpsOrg, options *AzureDevOpsOrgsClientBeginUpdateOptions) (*runtime.Poller[AzureDevOpsOrgsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, securityConnectorName, orgName, azureDevOpsOrg, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[AzureDevOpsOrgsClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[AzureDevOpsOrgsClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Update - Updates monitored Azure DevOps organization details.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01-preview
func (client *AzureDevOpsOrgsClient) update(ctx context.Context, resourceGroupName string, securityConnectorName string, orgName string, azureDevOpsOrg AzureDevOpsOrg, options *AzureDevOpsOrgsClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "AzureDevOpsOrgsClient.BeginUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, securityConnectorName, orgName, azureDevOpsOrg, options)
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
func (client *AzureDevOpsOrgsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, securityConnectorName string, orgName string, azureDevOpsOrg AzureDevOpsOrg, options *AzureDevOpsOrgsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/securityConnectors/{securityConnectorName}/devops/default/azureDevOpsOrgs/{orgName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if securityConnectorName == "" {
		return nil, errors.New("parameter securityConnectorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{securityConnectorName}", url.PathEscape(securityConnectorName))
	if orgName == "" {
		return nil, errors.New("parameter orgName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{orgName}", url.PathEscape(orgName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, azureDevOpsOrg); err != nil {
		return nil, err
	}
	return req, nil
}
