//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsecurityinsights

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

// ThreatIntelligenceIndicatorClient contains the methods for the ThreatIntelligenceIndicator group.
// Don't use this type directly, use NewThreatIntelligenceIndicatorClient() instead.
type ThreatIntelligenceIndicatorClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewThreatIntelligenceIndicatorClient creates a new instance of ThreatIntelligenceIndicatorClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewThreatIntelligenceIndicatorClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ThreatIntelligenceIndicatorClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ThreatIntelligenceIndicatorClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// AppendTags - Append tags to a threat intelligence indicator.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - name - Threat intelligence indicator name field.
//   - threatIntelligenceAppendTags - The threat intelligence append tags request body
//   - options - ThreatIntelligenceIndicatorClientAppendTagsOptions contains the optional parameters for the ThreatIntelligenceIndicatorClient.AppendTags
//     method.
func (client *ThreatIntelligenceIndicatorClient) AppendTags(ctx context.Context, resourceGroupName string, workspaceName string, name string, threatIntelligenceAppendTags ThreatIntelligenceAppendTags, options *ThreatIntelligenceIndicatorClientAppendTagsOptions) (ThreatIntelligenceIndicatorClientAppendTagsResponse, error) {
	var err error
	const operationName = "ThreatIntelligenceIndicatorClient.AppendTags"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.appendTagsCreateRequest(ctx, resourceGroupName, workspaceName, name, threatIntelligenceAppendTags, options)
	if err != nil {
		return ThreatIntelligenceIndicatorClientAppendTagsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ThreatIntelligenceIndicatorClientAppendTagsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ThreatIntelligenceIndicatorClientAppendTagsResponse{}, err
	}
	return ThreatIntelligenceIndicatorClientAppendTagsResponse{}, nil
}

// appendTagsCreateRequest creates the AppendTags request.
func (client *ThreatIntelligenceIndicatorClient) appendTagsCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, name string, threatIntelligenceAppendTags ThreatIntelligenceAppendTags, options *ThreatIntelligenceIndicatorClientAppendTagsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/threatIntelligence/main/indicators/{name}/appendTags"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, threatIntelligenceAppendTags); err != nil {
		return nil, err
	}
	return req, nil
}

// Create - Update a threat Intelligence indicator.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - name - Threat intelligence indicator name field.
//   - threatIntelligenceProperties - Properties of threat intelligence indicators to create and update.
//   - options - ThreatIntelligenceIndicatorClientCreateOptions contains the optional parameters for the ThreatIntelligenceIndicatorClient.Create
//     method.
func (client *ThreatIntelligenceIndicatorClient) Create(ctx context.Context, resourceGroupName string, workspaceName string, name string, threatIntelligenceProperties ThreatIntelligenceIndicatorModel, options *ThreatIntelligenceIndicatorClientCreateOptions) (ThreatIntelligenceIndicatorClientCreateResponse, error) {
	var err error
	const operationName = "ThreatIntelligenceIndicatorClient.Create"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createCreateRequest(ctx, resourceGroupName, workspaceName, name, threatIntelligenceProperties, options)
	if err != nil {
		return ThreatIntelligenceIndicatorClientCreateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ThreatIntelligenceIndicatorClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return ThreatIntelligenceIndicatorClientCreateResponse{}, err
	}
	resp, err := client.createHandleResponse(httpResp)
	return resp, err
}

// createCreateRequest creates the Create request.
func (client *ThreatIntelligenceIndicatorClient) createCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, name string, threatIntelligenceProperties ThreatIntelligenceIndicatorModel, options *ThreatIntelligenceIndicatorClientCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/threatIntelligence/main/indicators/{name}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, threatIntelligenceProperties); err != nil {
		return nil, err
	}
	return req, nil
}

// createHandleResponse handles the Create response.
func (client *ThreatIntelligenceIndicatorClient) createHandleResponse(resp *http.Response) (ThreatIntelligenceIndicatorClientCreateResponse, error) {
	result := ThreatIntelligenceIndicatorClientCreateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result); err != nil {
		return ThreatIntelligenceIndicatorClientCreateResponse{}, err
	}
	return result, nil
}

// CreateIndicator - Create a new threat intelligence indicator.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - threatIntelligenceProperties - Properties of threat intelligence indicators to create and update.
//   - options - ThreatIntelligenceIndicatorClientCreateIndicatorOptions contains the optional parameters for the ThreatIntelligenceIndicatorClient.CreateIndicator
//     method.
func (client *ThreatIntelligenceIndicatorClient) CreateIndicator(ctx context.Context, resourceGroupName string, workspaceName string, threatIntelligenceProperties ThreatIntelligenceIndicatorModel, options *ThreatIntelligenceIndicatorClientCreateIndicatorOptions) (ThreatIntelligenceIndicatorClientCreateIndicatorResponse, error) {
	var err error
	const operationName = "ThreatIntelligenceIndicatorClient.CreateIndicator"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createIndicatorCreateRequest(ctx, resourceGroupName, workspaceName, threatIntelligenceProperties, options)
	if err != nil {
		return ThreatIntelligenceIndicatorClientCreateIndicatorResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ThreatIntelligenceIndicatorClientCreateIndicatorResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return ThreatIntelligenceIndicatorClientCreateIndicatorResponse{}, err
	}
	resp, err := client.createIndicatorHandleResponse(httpResp)
	return resp, err
}

// createIndicatorCreateRequest creates the CreateIndicator request.
func (client *ThreatIntelligenceIndicatorClient) createIndicatorCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, threatIntelligenceProperties ThreatIntelligenceIndicatorModel, options *ThreatIntelligenceIndicatorClientCreateIndicatorOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/threatIntelligence/main/createIndicator"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, threatIntelligenceProperties); err != nil {
		return nil, err
	}
	return req, nil
}

// createIndicatorHandleResponse handles the CreateIndicator response.
func (client *ThreatIntelligenceIndicatorClient) createIndicatorHandleResponse(resp *http.Response) (ThreatIntelligenceIndicatorClientCreateIndicatorResponse, error) {
	result := ThreatIntelligenceIndicatorClientCreateIndicatorResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result); err != nil {
		return ThreatIntelligenceIndicatorClientCreateIndicatorResponse{}, err
	}
	return result, nil
}

// Delete - Delete a threat intelligence indicator.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - name - Threat intelligence indicator name field.
//   - options - ThreatIntelligenceIndicatorClientDeleteOptions contains the optional parameters for the ThreatIntelligenceIndicatorClient.Delete
//     method.
func (client *ThreatIntelligenceIndicatorClient) Delete(ctx context.Context, resourceGroupName string, workspaceName string, name string, options *ThreatIntelligenceIndicatorClientDeleteOptions) (ThreatIntelligenceIndicatorClientDeleteResponse, error) {
	var err error
	const operationName = "ThreatIntelligenceIndicatorClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, workspaceName, name, options)
	if err != nil {
		return ThreatIntelligenceIndicatorClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ThreatIntelligenceIndicatorClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ThreatIntelligenceIndicatorClientDeleteResponse{}, err
	}
	return ThreatIntelligenceIndicatorClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ThreatIntelligenceIndicatorClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, name string, options *ThreatIntelligenceIndicatorClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/threatIntelligence/main/indicators/{name}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - View a threat intelligence indicator by name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - name - Threat intelligence indicator name field.
//   - options - ThreatIntelligenceIndicatorClientGetOptions contains the optional parameters for the ThreatIntelligenceIndicatorClient.Get
//     method.
func (client *ThreatIntelligenceIndicatorClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, name string, options *ThreatIntelligenceIndicatorClientGetOptions) (ThreatIntelligenceIndicatorClientGetResponse, error) {
	var err error
	const operationName = "ThreatIntelligenceIndicatorClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, name, options)
	if err != nil {
		return ThreatIntelligenceIndicatorClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ThreatIntelligenceIndicatorClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ThreatIntelligenceIndicatorClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ThreatIntelligenceIndicatorClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, name string, options *ThreatIntelligenceIndicatorClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/threatIntelligence/main/indicators/{name}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ThreatIntelligenceIndicatorClient) getHandleResponse(resp *http.Response) (ThreatIntelligenceIndicatorClientGetResponse, error) {
	result := ThreatIntelligenceIndicatorClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result); err != nil {
		return ThreatIntelligenceIndicatorClientGetResponse{}, err
	}
	return result, nil
}

// NewQueryIndicatorsPager - Query threat intelligence indicators as per filtering criteria.
//
// Generated from API version 2024-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - threatIntelligenceFilteringCriteria - Filtering criteria for querying threat intelligence indicators.
//   - options - ThreatIntelligenceIndicatorClientQueryIndicatorsOptions contains the optional parameters for the ThreatIntelligenceIndicatorClient.NewQueryIndicatorsPager
//     method.
func (client *ThreatIntelligenceIndicatorClient) NewQueryIndicatorsPager(resourceGroupName string, workspaceName string, threatIntelligenceFilteringCriteria ThreatIntelligenceFilteringCriteria, options *ThreatIntelligenceIndicatorClientQueryIndicatorsOptions) *runtime.Pager[ThreatIntelligenceIndicatorClientQueryIndicatorsResponse] {
	return runtime.NewPager(runtime.PagingHandler[ThreatIntelligenceIndicatorClientQueryIndicatorsResponse]{
		More: func(page ThreatIntelligenceIndicatorClientQueryIndicatorsResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ThreatIntelligenceIndicatorClientQueryIndicatorsResponse) (ThreatIntelligenceIndicatorClientQueryIndicatorsResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ThreatIntelligenceIndicatorClient.NewQueryIndicatorsPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.queryIndicatorsCreateRequest(ctx, resourceGroupName, workspaceName, threatIntelligenceFilteringCriteria, options)
			}, nil)
			if err != nil {
				return ThreatIntelligenceIndicatorClientQueryIndicatorsResponse{}, err
			}
			return client.queryIndicatorsHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// queryIndicatorsCreateRequest creates the QueryIndicators request.
func (client *ThreatIntelligenceIndicatorClient) queryIndicatorsCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, threatIntelligenceFilteringCriteria ThreatIntelligenceFilteringCriteria, options *ThreatIntelligenceIndicatorClientQueryIndicatorsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/threatIntelligence/main/queryIndicators"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, threatIntelligenceFilteringCriteria); err != nil {
		return nil, err
	}
	return req, nil
}

// queryIndicatorsHandleResponse handles the QueryIndicators response.
func (client *ThreatIntelligenceIndicatorClient) queryIndicatorsHandleResponse(resp *http.Response) (ThreatIntelligenceIndicatorClientQueryIndicatorsResponse, error) {
	result := ThreatIntelligenceIndicatorClientQueryIndicatorsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ThreatIntelligenceInformationList); err != nil {
		return ThreatIntelligenceIndicatorClientQueryIndicatorsResponse{}, err
	}
	return result, nil
}

// ReplaceTags - Replace tags added to a threat intelligence indicator.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - name - Threat intelligence indicator name field.
//   - threatIntelligenceReplaceTags - Tags in the threat intelligence indicator to be replaced.
//   - options - ThreatIntelligenceIndicatorClientReplaceTagsOptions contains the optional parameters for the ThreatIntelligenceIndicatorClient.ReplaceTags
//     method.
func (client *ThreatIntelligenceIndicatorClient) ReplaceTags(ctx context.Context, resourceGroupName string, workspaceName string, name string, threatIntelligenceReplaceTags ThreatIntelligenceIndicatorModel, options *ThreatIntelligenceIndicatorClientReplaceTagsOptions) (ThreatIntelligenceIndicatorClientReplaceTagsResponse, error) {
	var err error
	const operationName = "ThreatIntelligenceIndicatorClient.ReplaceTags"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.replaceTagsCreateRequest(ctx, resourceGroupName, workspaceName, name, threatIntelligenceReplaceTags, options)
	if err != nil {
		return ThreatIntelligenceIndicatorClientReplaceTagsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ThreatIntelligenceIndicatorClientReplaceTagsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ThreatIntelligenceIndicatorClientReplaceTagsResponse{}, err
	}
	resp, err := client.replaceTagsHandleResponse(httpResp)
	return resp, err
}

// replaceTagsCreateRequest creates the ReplaceTags request.
func (client *ThreatIntelligenceIndicatorClient) replaceTagsCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, name string, threatIntelligenceReplaceTags ThreatIntelligenceIndicatorModel, options *ThreatIntelligenceIndicatorClientReplaceTagsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/threatIntelligence/main/indicators/{name}/replaceTags"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, threatIntelligenceReplaceTags); err != nil {
		return nil, err
	}
	return req, nil
}

// replaceTagsHandleResponse handles the ReplaceTags response.
func (client *ThreatIntelligenceIndicatorClient) replaceTagsHandleResponse(resp *http.Response) (ThreatIntelligenceIndicatorClientReplaceTagsResponse, error) {
	result := ThreatIntelligenceIndicatorClientReplaceTagsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result); err != nil {
		return ThreatIntelligenceIndicatorClientReplaceTagsResponse{}, err
	}
	return result, nil
}
