//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armappplatform

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

// APIPortalsClient contains the methods for the APIPortals group.
// Don't use this type directly, use NewAPIPortalsClient() instead.
type APIPortalsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewAPIPortalsClient creates a new instance of APIPortalsClient with the specified values.
//   - subscriptionID - Gets subscription ID which uniquely identify the Microsoft Azure subscription. The subscription ID forms
//     part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewAPIPortalsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*APIPortalsClient, error) {
	cl, err := arm.NewClient(moduleName+".APIPortalsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &APIPortalsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create the default API portal or update the existing API portal.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serviceName - The name of the Service resource.
//   - apiPortalName - The name of API portal.
//   - apiPortalResource - The API portal for the create or update operation
//   - options - APIPortalsClientBeginCreateOrUpdateOptions contains the optional parameters for the APIPortalsClient.BeginCreateOrUpdate
//     method.
func (client *APIPortalsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, apiPortalName string, apiPortalResource APIPortalResource, options *APIPortalsClientBeginCreateOrUpdateOptions) (*runtime.Poller[APIPortalsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, serviceName, apiPortalName, apiPortalResource, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[APIPortalsClientCreateOrUpdateResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[APIPortalsClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Create the default API portal or update the existing API portal.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01-preview
func (client *APIPortalsClient) createOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, apiPortalName string, apiPortalResource APIPortalResource, options *APIPortalsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serviceName, apiPortalName, apiPortalResource, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *APIPortalsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, apiPortalName string, apiPortalResource APIPortalResource, options *APIPortalsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/apiPortals/{apiPortalName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if apiPortalName == "" {
		return nil, errors.New("parameter apiPortalName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiPortalName}", url.PathEscape(apiPortalName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, apiPortalResource)
}

// BeginDelete - Delete the default API portal.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serviceName - The name of the Service resource.
//   - apiPortalName - The name of API portal.
//   - options - APIPortalsClientBeginDeleteOptions contains the optional parameters for the APIPortalsClient.BeginDelete method.
func (client *APIPortalsClient) BeginDelete(ctx context.Context, resourceGroupName string, serviceName string, apiPortalName string, options *APIPortalsClientBeginDeleteOptions) (*runtime.Poller[APIPortalsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, serviceName, apiPortalName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[APIPortalsClientDeleteResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[APIPortalsClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Delete the default API portal.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01-preview
func (client *APIPortalsClient) deleteOperation(ctx context.Context, resourceGroupName string, serviceName string, apiPortalName string, options *APIPortalsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serviceName, apiPortalName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *APIPortalsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, apiPortalName string, options *APIPortalsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/apiPortals/{apiPortalName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if apiPortalName == "" {
		return nil, errors.New("parameter apiPortalName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiPortalName}", url.PathEscape(apiPortalName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get the API portal and its properties.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serviceName - The name of the Service resource.
//   - apiPortalName - The name of API portal.
//   - options - APIPortalsClientGetOptions contains the optional parameters for the APIPortalsClient.Get method.
func (client *APIPortalsClient) Get(ctx context.Context, resourceGroupName string, serviceName string, apiPortalName string, options *APIPortalsClientGetOptions) (APIPortalsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serviceName, apiPortalName, options)
	if err != nil {
		return APIPortalsClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return APIPortalsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return APIPortalsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *APIPortalsClient) getCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, apiPortalName string, options *APIPortalsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/apiPortals/{apiPortalName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if apiPortalName == "" {
		return nil, errors.New("parameter apiPortalName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiPortalName}", url.PathEscape(apiPortalName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *APIPortalsClient) getHandleResponse(resp *http.Response) (APIPortalsClientGetResponse, error) {
	result := APIPortalsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.APIPortalResource); err != nil {
		return APIPortalsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Handles requests to list all resources in a Service.
//
// Generated from API version 2023-03-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serviceName - The name of the Service resource.
//   - options - APIPortalsClientListOptions contains the optional parameters for the APIPortalsClient.NewListPager method.
func (client *APIPortalsClient) NewListPager(resourceGroupName string, serviceName string, options *APIPortalsClientListOptions) *runtime.Pager[APIPortalsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[APIPortalsClientListResponse]{
		More: func(page APIPortalsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *APIPortalsClientListResponse) (APIPortalsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, serviceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return APIPortalsClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return APIPortalsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return APIPortalsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *APIPortalsClient) listCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, options *APIPortalsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/apiPortals"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *APIPortalsClient) listHandleResponse(resp *http.Response) (APIPortalsClientListResponse, error) {
	result := APIPortalsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.APIPortalResourceCollection); err != nil {
		return APIPortalsClientListResponse{}, err
	}
	return result, nil
}

// ValidateDomain - Check the domains are valid as well as not in use.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serviceName - The name of the Service resource.
//   - apiPortalName - The name of API portal.
//   - validatePayload - Custom domain payload to be validated
//   - options - APIPortalsClientValidateDomainOptions contains the optional parameters for the APIPortalsClient.ValidateDomain
//     method.
func (client *APIPortalsClient) ValidateDomain(ctx context.Context, resourceGroupName string, serviceName string, apiPortalName string, validatePayload CustomDomainValidatePayload, options *APIPortalsClientValidateDomainOptions) (APIPortalsClientValidateDomainResponse, error) {
	req, err := client.validateDomainCreateRequest(ctx, resourceGroupName, serviceName, apiPortalName, validatePayload, options)
	if err != nil {
		return APIPortalsClientValidateDomainResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return APIPortalsClientValidateDomainResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return APIPortalsClientValidateDomainResponse{}, runtime.NewResponseError(resp)
	}
	return client.validateDomainHandleResponse(resp)
}

// validateDomainCreateRequest creates the ValidateDomain request.
func (client *APIPortalsClient) validateDomainCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, apiPortalName string, validatePayload CustomDomainValidatePayload, options *APIPortalsClientValidateDomainOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/apiPortals/{apiPortalName}/validateDomain"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if apiPortalName == "" {
		return nil, errors.New("parameter apiPortalName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiPortalName}", url.PathEscape(apiPortalName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, validatePayload)
}

// validateDomainHandleResponse handles the ValidateDomain response.
func (client *APIPortalsClient) validateDomainHandleResponse(resp *http.Response) (APIPortalsClientValidateDomainResponse, error) {
	result := APIPortalsClientValidateDomainResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CustomDomainValidateResult); err != nil {
		return APIPortalsClientValidateDomainResponse{}, err
	}
	return result, nil
}
