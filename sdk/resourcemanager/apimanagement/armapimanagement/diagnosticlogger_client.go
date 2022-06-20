//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armapimanagement

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
	"strconv"
	"strings"
)

// DiagnosticLoggerClient contains the methods for the DiagnosticLogger group.
// Don't use this type directly, use NewDiagnosticLoggerClient() instead.
type DiagnosticLoggerClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewDiagnosticLoggerClient creates a new instance of DiagnosticLoggerClient with the specified values.
// subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewDiagnosticLoggerClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DiagnosticLoggerClient, error) {
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
	client := &DiagnosticLoggerClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CheckEntityExists - Checks that logger entity specified by identifier is associated with the diagnostics entity.
// Generated from API version 2022-01-09
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// diagnosticID - Diagnostic identifier. Must be unique in the current API Management service instance.
// loggerid - Logger identifier. Must be unique in the API Management service instance.
// options - DiagnosticLoggerClientCheckEntityExistsOptions contains the optional parameters for the DiagnosticLoggerClient.CheckEntityExists
// method.
func (client *DiagnosticLoggerClient) CheckEntityExists(ctx context.Context, resourceGroupName string, serviceName string, diagnosticID string, loggerid string, options *DiagnosticLoggerClientCheckEntityExistsOptions) (DiagnosticLoggerClientCheckEntityExistsResponse, error) {
	req, err := client.checkEntityExistsCreateRequest(ctx, resourceGroupName, serviceName, diagnosticID, loggerid, options)
	if err != nil {
		return DiagnosticLoggerClientCheckEntityExistsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DiagnosticLoggerClientCheckEntityExistsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusNoContent) {
		return DiagnosticLoggerClientCheckEntityExistsResponse{}, runtime.NewResponseError(resp)
	}
	return DiagnosticLoggerClientCheckEntityExistsResponse{Success: resp.StatusCode >= 200 && resp.StatusCode < 300}, nil
}

// checkEntityExistsCreateRequest creates the CheckEntityExists request.
func (client *DiagnosticLoggerClient) checkEntityExistsCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, diagnosticID string, loggerid string, options *DiagnosticLoggerClientCheckEntityExistsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/diagnostics/{diagnosticId}/loggers/{loggerid}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if diagnosticID == "" {
		return nil, errors.New("parameter diagnosticID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{diagnosticId}", url.PathEscape(diagnosticID))
	if loggerid == "" {
		return nil, errors.New("parameter loggerid cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{loggerid}", url.PathEscape(loggerid))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodHead, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-09")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// CreateOrUpdate - Attaches a logger to a diagnostic.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-09
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// diagnosticID - Diagnostic identifier. Must be unique in the current API Management service instance.
// loggerid - Logger identifier. Must be unique in the API Management service instance.
// options - DiagnosticLoggerClientCreateOrUpdateOptions contains the optional parameters for the DiagnosticLoggerClient.CreateOrUpdate
// method.
func (client *DiagnosticLoggerClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, diagnosticID string, loggerid string, options *DiagnosticLoggerClientCreateOrUpdateOptions) (DiagnosticLoggerClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serviceName, diagnosticID, loggerid, options)
	if err != nil {
		return DiagnosticLoggerClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DiagnosticLoggerClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return DiagnosticLoggerClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *DiagnosticLoggerClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, diagnosticID string, loggerid string, options *DiagnosticLoggerClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/diagnostics/{diagnosticId}/loggers/{loggerid}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if diagnosticID == "" {
		return nil, errors.New("parameter diagnosticID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{diagnosticId}", url.PathEscape(diagnosticID))
	if loggerid == "" {
		return nil, errors.New("parameter loggerid cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{loggerid}", url.PathEscape(loggerid))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-09")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *DiagnosticLoggerClient) createOrUpdateHandleResponse(resp *http.Response) (DiagnosticLoggerClientCreateOrUpdateResponse, error) {
	result := DiagnosticLoggerClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LoggerContract); err != nil {
		return DiagnosticLoggerClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes the specified Logger from Diagnostic.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-09
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// diagnosticID - Diagnostic identifier. Must be unique in the current API Management service instance.
// loggerid - Logger identifier. Must be unique in the API Management service instance.
// options - DiagnosticLoggerClientDeleteOptions contains the optional parameters for the DiagnosticLoggerClient.Delete method.
func (client *DiagnosticLoggerClient) Delete(ctx context.Context, resourceGroupName string, serviceName string, diagnosticID string, loggerid string, options *DiagnosticLoggerClientDeleteOptions) (DiagnosticLoggerClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serviceName, diagnosticID, loggerid, options)
	if err != nil {
		return DiagnosticLoggerClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DiagnosticLoggerClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusNoContent) {
		return DiagnosticLoggerClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return DiagnosticLoggerClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *DiagnosticLoggerClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, diagnosticID string, loggerid string, options *DiagnosticLoggerClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/diagnostics/{diagnosticId}/loggers/{loggerid}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if diagnosticID == "" {
		return nil, errors.New("parameter diagnosticID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{diagnosticId}", url.PathEscape(diagnosticID))
	if loggerid == "" {
		return nil, errors.New("parameter loggerid cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{loggerid}", url.PathEscape(loggerid))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-09")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// NewListByServicePager - Lists all loggers associated with the specified Diagnostic of the API Management service instance.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-09
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// diagnosticID - Diagnostic identifier. Must be unique in the current API Management service instance.
// options - DiagnosticLoggerClientListByServiceOptions contains the optional parameters for the DiagnosticLoggerClient.ListByService
// method.
func (client *DiagnosticLoggerClient) NewListByServicePager(resourceGroupName string, serviceName string, diagnosticID string, options *DiagnosticLoggerClientListByServiceOptions) *runtime.Pager[DiagnosticLoggerClientListByServiceResponse] {
	return runtime.NewPager(runtime.PagingHandler[DiagnosticLoggerClientListByServiceResponse]{
		More: func(page DiagnosticLoggerClientListByServiceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DiagnosticLoggerClientListByServiceResponse) (DiagnosticLoggerClientListByServiceResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByServiceCreateRequest(ctx, resourceGroupName, serviceName, diagnosticID, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DiagnosticLoggerClientListByServiceResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return DiagnosticLoggerClientListByServiceResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DiagnosticLoggerClientListByServiceResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByServiceHandleResponse(resp)
		},
	})
}

// listByServiceCreateRequest creates the ListByService request.
func (client *DiagnosticLoggerClient) listByServiceCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, diagnosticID string, options *DiagnosticLoggerClientListByServiceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/diagnostics/{diagnosticId}/loggers"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if diagnosticID == "" {
		return nil, errors.New("parameter diagnosticID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{diagnosticId}", url.PathEscape(diagnosticID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-09")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", strconv.FormatInt(int64(*options.Skip), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByServiceHandleResponse handles the ListByService response.
func (client *DiagnosticLoggerClient) listByServiceHandleResponse(resp *http.Response) (DiagnosticLoggerClientListByServiceResponse, error) {
	result := DiagnosticLoggerClientListByServiceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LoggerCollection); err != nil {
		return DiagnosticLoggerClientListByServiceResponse{}, err
	}
	return result, nil
}
