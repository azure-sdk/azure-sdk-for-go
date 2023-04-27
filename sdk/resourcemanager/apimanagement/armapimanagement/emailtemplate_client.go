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

// EmailTemplateClient contains the methods for the EmailTemplate group.
// Don't use this type directly, use NewEmailTemplateClient() instead.
type EmailTemplateClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewEmailTemplateClient creates a new instance of EmailTemplateClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewEmailTemplateClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*EmailTemplateClient, error) {
	cl, err := arm.NewClient(moduleName+".EmailTemplateClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &EmailTemplateClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Updates an Email Template.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - templateName - Email Template Name Identifier.
//   - parameters - Email Template update parameters.
//   - options - EmailTemplateClientCreateOrUpdateOptions contains the optional parameters for the EmailTemplateClient.CreateOrUpdate
//     method.
func (client *EmailTemplateClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, templateName TemplateName, parameters EmailTemplateUpdateParameters, options *EmailTemplateClientCreateOrUpdateOptions) (EmailTemplateClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serviceName, templateName, parameters, options)
	if err != nil {
		return EmailTemplateClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return EmailTemplateClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return EmailTemplateClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *EmailTemplateClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, templateName TemplateName, parameters EmailTemplateUpdateParameters, options *EmailTemplateClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/templates/{templateName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if templateName == "" {
		return nil, errors.New("parameter templateName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{templateName}", url.PathEscape(string(templateName)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header["If-Match"] = []string{*options.IfMatch}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *EmailTemplateClient) createOrUpdateHandleResponse(resp *http.Response) (EmailTemplateClientCreateOrUpdateResponse, error) {
	result := EmailTemplateClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EmailTemplateContract); err != nil {
		return EmailTemplateClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Reset the Email Template to default template provided by the API Management service instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - templateName - Email Template Name Identifier.
//   - ifMatch - ETag of the Entity. ETag should match the current entity state from the header response of the GET request or
//     it should be * for unconditional update.
//   - options - EmailTemplateClientDeleteOptions contains the optional parameters for the EmailTemplateClient.Delete method.
func (client *EmailTemplateClient) Delete(ctx context.Context, resourceGroupName string, serviceName string, templateName TemplateName, ifMatch string, options *EmailTemplateClientDeleteOptions) (EmailTemplateClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serviceName, templateName, ifMatch, options)
	if err != nil {
		return EmailTemplateClientDeleteResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return EmailTemplateClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return EmailTemplateClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return EmailTemplateClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *EmailTemplateClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, templateName TemplateName, ifMatch string, options *EmailTemplateClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/templates/{templateName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if templateName == "" {
		return nil, errors.New("parameter templateName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{templateName}", url.PathEscape(string(templateName)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["If-Match"] = []string{ifMatch}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the details of the email template specified by its identifier.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - templateName - Email Template Name Identifier.
//   - options - EmailTemplateClientGetOptions contains the optional parameters for the EmailTemplateClient.Get method.
func (client *EmailTemplateClient) Get(ctx context.Context, resourceGroupName string, serviceName string, templateName TemplateName, options *EmailTemplateClientGetOptions) (EmailTemplateClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serviceName, templateName, options)
	if err != nil {
		return EmailTemplateClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return EmailTemplateClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return EmailTemplateClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *EmailTemplateClient) getCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, templateName TemplateName, options *EmailTemplateClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/templates/{templateName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if templateName == "" {
		return nil, errors.New("parameter templateName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{templateName}", url.PathEscape(string(templateName)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *EmailTemplateClient) getHandleResponse(resp *http.Response) (EmailTemplateClientGetResponse, error) {
	result := EmailTemplateClientGetResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.EmailTemplateContract); err != nil {
		return EmailTemplateClientGetResponse{}, err
	}
	return result, nil
}

// GetEntityTag - Gets the entity state (Etag) version of the email template specified by its identifier.
//
// Generated from API version 2022-09-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - templateName - Email Template Name Identifier.
//   - options - EmailTemplateClientGetEntityTagOptions contains the optional parameters for the EmailTemplateClient.GetEntityTag
//     method.
func (client *EmailTemplateClient) GetEntityTag(ctx context.Context, resourceGroupName string, serviceName string, templateName TemplateName, options *EmailTemplateClientGetEntityTagOptions) (EmailTemplateClientGetEntityTagResponse, error) {
	req, err := client.getEntityTagCreateRequest(ctx, resourceGroupName, serviceName, templateName, options)
	if err != nil {
		return EmailTemplateClientGetEntityTagResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return EmailTemplateClientGetEntityTagResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return EmailTemplateClientGetEntityTagResponse{}, runtime.NewResponseError(resp)
	}
	return client.getEntityTagHandleResponse(resp)
}

// getEntityTagCreateRequest creates the GetEntityTag request.
func (client *EmailTemplateClient) getEntityTagCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, templateName TemplateName, options *EmailTemplateClientGetEntityTagOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/templates/{templateName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if templateName == "" {
		return nil, errors.New("parameter templateName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{templateName}", url.PathEscape(string(templateName)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodHead, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getEntityTagHandleResponse handles the GetEntityTag response.
func (client *EmailTemplateClient) getEntityTagHandleResponse(resp *http.Response) (EmailTemplateClientGetEntityTagResponse, error) {
	result := EmailTemplateClientGetEntityTagResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	result.Success = resp.StatusCode >= 200 && resp.StatusCode < 300
	return result, nil
}

// NewListByServicePager - Gets all email templates
//
// Generated from API version 2022-09-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - options - EmailTemplateClientListByServiceOptions contains the optional parameters for the EmailTemplateClient.NewListByServicePager
//     method.
func (client *EmailTemplateClient) NewListByServicePager(resourceGroupName string, serviceName string, options *EmailTemplateClientListByServiceOptions) *runtime.Pager[EmailTemplateClientListByServiceResponse] {
	return runtime.NewPager(runtime.PagingHandler[EmailTemplateClientListByServiceResponse]{
		More: func(page EmailTemplateClientListByServiceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *EmailTemplateClientListByServiceResponse) (EmailTemplateClientListByServiceResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByServiceCreateRequest(ctx, resourceGroupName, serviceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return EmailTemplateClientListByServiceResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return EmailTemplateClientListByServiceResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return EmailTemplateClientListByServiceResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByServiceHandleResponse(resp)
		},
	})
}

// listByServiceCreateRequest creates the ListByService request.
func (client *EmailTemplateClient) listByServiceCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, options *EmailTemplateClientListByServiceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/templates"
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
	reqQP.Set("api-version", "2022-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByServiceHandleResponse handles the ListByService response.
func (client *EmailTemplateClient) listByServiceHandleResponse(resp *http.Response) (EmailTemplateClientListByServiceResponse, error) {
	result := EmailTemplateClientListByServiceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EmailTemplateCollection); err != nil {
		return EmailTemplateClientListByServiceResponse{}, err
	}
	return result, nil
}

// Update - Updates API Management email template
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - templateName - Email Template Name Identifier.
//   - ifMatch - ETag of the Entity. ETag should match the current entity state from the header response of the GET request or
//     it should be * for unconditional update.
//   - parameters - Update parameters.
//   - options - EmailTemplateClientUpdateOptions contains the optional parameters for the EmailTemplateClient.Update method.
func (client *EmailTemplateClient) Update(ctx context.Context, resourceGroupName string, serviceName string, templateName TemplateName, ifMatch string, parameters EmailTemplateUpdateParameters, options *EmailTemplateClientUpdateOptions) (EmailTemplateClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, serviceName, templateName, ifMatch, parameters, options)
	if err != nil {
		return EmailTemplateClientUpdateResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return EmailTemplateClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return EmailTemplateClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *EmailTemplateClient) updateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, templateName TemplateName, ifMatch string, parameters EmailTemplateUpdateParameters, options *EmailTemplateClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/templates/{templateName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if templateName == "" {
		return nil, errors.New("parameter templateName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{templateName}", url.PathEscape(string(templateName)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["If-Match"] = []string{ifMatch}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateHandleResponse handles the Update response.
func (client *EmailTemplateClient) updateHandleResponse(resp *http.Response) (EmailTemplateClientUpdateResponse, error) {
	result := EmailTemplateClientUpdateResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.EmailTemplateContract); err != nil {
		return EmailTemplateClientUpdateResponse{}, err
	}
	return result, nil
}
