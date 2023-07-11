//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armapimanagement

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

// APIDiagnosticClient contains the methods for the APIDiagnostic group.
// Don't use this type directly, use NewAPIDiagnosticClient() instead.
type APIDiagnosticClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewAPIDiagnosticClient creates a new instance of APIDiagnosticClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewAPIDiagnosticClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*APIDiagnosticClient, error) {
	cl, err := arm.NewClient(moduleName+".APIDiagnosticClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &APIDiagnosticClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Creates a new Diagnostic for an API or updates an existing one.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - apiID - API identifier. Must be unique in the current API Management service instance.
//   - diagnosticID - Diagnostic identifier. Must be unique in the current API Management service instance.
//   - parameters - Create parameters.
//   - options - APIDiagnosticClientCreateOrUpdateOptions contains the optional parameters for the APIDiagnosticClient.CreateOrUpdate
//     method.
func (client *APIDiagnosticClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, apiID string, diagnosticID string, parameters DiagnosticContract, options *APIDiagnosticClientCreateOrUpdateOptions) (APIDiagnosticClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serviceName, apiID, diagnosticID, parameters, options)
	if err != nil {
		return APIDiagnosticClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return APIDiagnosticClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return APIDiagnosticClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *APIDiagnosticClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, apiID string, diagnosticID string, parameters DiagnosticContract, options *APIDiagnosticClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/diagnostics/{diagnosticId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if apiID == "" {
		return nil, errors.New("parameter apiID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiId}", url.PathEscape(apiID))
	if diagnosticID == "" {
		return nil, errors.New("parameter diagnosticID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{diagnosticId}", url.PathEscape(diagnosticID))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header["If-Match"] = []string{*options.IfMatch}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *APIDiagnosticClient) createOrUpdateHandleResponse(resp *http.Response) (APIDiagnosticClientCreateOrUpdateResponse, error) {
	result := APIDiagnosticClientCreateOrUpdateResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.DiagnosticContract); err != nil {
		return APIDiagnosticClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes the specified Diagnostic from an API.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - apiID - API identifier. Must be unique in the current API Management service instance.
//   - diagnosticID - Diagnostic identifier. Must be unique in the current API Management service instance.
//   - ifMatch - ETag of the Entity. ETag should match the current entity state from the header response of the GET request or
//     it should be * for unconditional update.
//   - options - APIDiagnosticClientDeleteOptions contains the optional parameters for the APIDiagnosticClient.Delete method.
func (client *APIDiagnosticClient) Delete(ctx context.Context, resourceGroupName string, serviceName string, apiID string, diagnosticID string, ifMatch string, options *APIDiagnosticClientDeleteOptions) (APIDiagnosticClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serviceName, apiID, diagnosticID, ifMatch, options)
	if err != nil {
		return APIDiagnosticClientDeleteResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return APIDiagnosticClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return APIDiagnosticClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return APIDiagnosticClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *APIDiagnosticClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, apiID string, diagnosticID string, ifMatch string, options *APIDiagnosticClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/diagnostics/{diagnosticId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if apiID == "" {
		return nil, errors.New("parameter apiID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiId}", url.PathEscape(apiID))
	if diagnosticID == "" {
		return nil, errors.New("parameter diagnosticID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{diagnosticId}", url.PathEscape(diagnosticID))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["If-Match"] = []string{ifMatch}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the details of the Diagnostic for an API specified by its identifier.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - apiID - API identifier. Must be unique in the current API Management service instance.
//   - diagnosticID - Diagnostic identifier. Must be unique in the current API Management service instance.
//   - options - APIDiagnosticClientGetOptions contains the optional parameters for the APIDiagnosticClient.Get method.
func (client *APIDiagnosticClient) Get(ctx context.Context, resourceGroupName string, serviceName string, apiID string, diagnosticID string, options *APIDiagnosticClientGetOptions) (APIDiagnosticClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serviceName, apiID, diagnosticID, options)
	if err != nil {
		return APIDiagnosticClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return APIDiagnosticClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return APIDiagnosticClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *APIDiagnosticClient) getCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, apiID string, diagnosticID string, options *APIDiagnosticClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/diagnostics/{diagnosticId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if apiID == "" {
		return nil, errors.New("parameter apiID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiId}", url.PathEscape(apiID))
	if diagnosticID == "" {
		return nil, errors.New("parameter diagnosticID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{diagnosticId}", url.PathEscape(diagnosticID))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *APIDiagnosticClient) getHandleResponse(resp *http.Response) (APIDiagnosticClientGetResponse, error) {
	result := APIDiagnosticClientGetResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.DiagnosticContract); err != nil {
		return APIDiagnosticClientGetResponse{}, err
	}
	return result, nil
}

// GetEntityTag - Gets the entity state (Etag) version of the Diagnostic for an API specified by its identifier.
//
// Generated from API version 2023-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - apiID - API identifier. Must be unique in the current API Management service instance.
//   - diagnosticID - Diagnostic identifier. Must be unique in the current API Management service instance.
//   - options - APIDiagnosticClientGetEntityTagOptions contains the optional parameters for the APIDiagnosticClient.GetEntityTag
//     method.
func (client *APIDiagnosticClient) GetEntityTag(ctx context.Context, resourceGroupName string, serviceName string, apiID string, diagnosticID string, options *APIDiagnosticClientGetEntityTagOptions) (APIDiagnosticClientGetEntityTagResponse, error) {
	req, err := client.getEntityTagCreateRequest(ctx, resourceGroupName, serviceName, apiID, diagnosticID, options)
	if err != nil {
		return APIDiagnosticClientGetEntityTagResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return APIDiagnosticClientGetEntityTagResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return APIDiagnosticClientGetEntityTagResponse{}, runtime.NewResponseError(resp)
	}
	return client.getEntityTagHandleResponse(resp)
}

// getEntityTagCreateRequest creates the GetEntityTag request.
func (client *APIDiagnosticClient) getEntityTagCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, apiID string, diagnosticID string, options *APIDiagnosticClientGetEntityTagOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/diagnostics/{diagnosticId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if apiID == "" {
		return nil, errors.New("parameter apiID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiId}", url.PathEscape(apiID))
	if diagnosticID == "" {
		return nil, errors.New("parameter diagnosticID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{diagnosticId}", url.PathEscape(diagnosticID))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodHead, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getEntityTagHandleResponse handles the GetEntityTag response.
func (client *APIDiagnosticClient) getEntityTagHandleResponse(resp *http.Response) (APIDiagnosticClientGetEntityTagResponse, error) {
	result := APIDiagnosticClientGetEntityTagResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	result.Success = resp.StatusCode >= 200 && resp.StatusCode < 300
	return result, nil
}

// NewListByServicePager - Lists all diagnostics of an API.
//
// Generated from API version 2023-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - apiID - API identifier. Must be unique in the current API Management service instance.
//   - options - APIDiagnosticClientListByServiceOptions contains the optional parameters for the APIDiagnosticClient.NewListByServicePager
//     method.
func (client *APIDiagnosticClient) NewListByServicePager(resourceGroupName string, serviceName string, apiID string, options *APIDiagnosticClientListByServiceOptions) *runtime.Pager[APIDiagnosticClientListByServiceResponse] {
	return runtime.NewPager(runtime.PagingHandler[APIDiagnosticClientListByServiceResponse]{
		More: func(page APIDiagnosticClientListByServiceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *APIDiagnosticClientListByServiceResponse) (APIDiagnosticClientListByServiceResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByServiceCreateRequest(ctx, resourceGroupName, serviceName, apiID, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return APIDiagnosticClientListByServiceResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return APIDiagnosticClientListByServiceResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return APIDiagnosticClientListByServiceResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByServiceHandleResponse(resp)
		},
	})
}

// listByServiceCreateRequest creates the ListByService request.
func (client *APIDiagnosticClient) listByServiceCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, apiID string, options *APIDiagnosticClientListByServiceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/diagnostics"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if apiID == "" {
		return nil, errors.New("parameter apiID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiId}", url.PathEscape(apiID))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
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
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", strconv.FormatInt(int64(*options.Skip), 10))
	}
	reqQP.Set("api-version", "2023-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByServiceHandleResponse handles the ListByService response.
func (client *APIDiagnosticClient) listByServiceHandleResponse(resp *http.Response) (APIDiagnosticClientListByServiceResponse, error) {
	result := APIDiagnosticClientListByServiceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DiagnosticCollection); err != nil {
		return APIDiagnosticClientListByServiceResponse{}, err
	}
	return result, nil
}

// Update - Updates the details of the Diagnostic for an API specified by its identifier.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - apiID - API identifier. Must be unique in the current API Management service instance.
//   - diagnosticID - Diagnostic identifier. Must be unique in the current API Management service instance.
//   - ifMatch - ETag of the Entity. ETag should match the current entity state from the header response of the GET request or
//     it should be * for unconditional update.
//   - parameters - Diagnostic Update parameters.
//   - options - APIDiagnosticClientUpdateOptions contains the optional parameters for the APIDiagnosticClient.Update method.
func (client *APIDiagnosticClient) Update(ctx context.Context, resourceGroupName string, serviceName string, apiID string, diagnosticID string, ifMatch string, parameters DiagnosticContract, options *APIDiagnosticClientUpdateOptions) (APIDiagnosticClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, serviceName, apiID, diagnosticID, ifMatch, parameters, options)
	if err != nil {
		return APIDiagnosticClientUpdateResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return APIDiagnosticClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return APIDiagnosticClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *APIDiagnosticClient) updateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, apiID string, diagnosticID string, ifMatch string, parameters DiagnosticContract, options *APIDiagnosticClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/diagnostics/{diagnosticId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if apiID == "" {
		return nil, errors.New("parameter apiID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiId}", url.PathEscape(apiID))
	if diagnosticID == "" {
		return nil, errors.New("parameter diagnosticID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{diagnosticId}", url.PathEscape(diagnosticID))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["If-Match"] = []string{ifMatch}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateHandleResponse handles the Update response.
func (client *APIDiagnosticClient) updateHandleResponse(resp *http.Response) (APIDiagnosticClientUpdateResponse, error) {
	result := APIDiagnosticClientUpdateResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.DiagnosticContract); err != nil {
		return APIDiagnosticClientUpdateResponse{}, err
	}
	return result, nil
}
