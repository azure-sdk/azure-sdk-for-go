//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armappservice

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

// CertificatesClient contains the methods for the Certificates group.
// Don't use this type directly, use NewCertificatesClient() instead.
type CertificatesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewCertificatesClient creates a new instance of CertificatesClient with the specified values.
//   - subscriptionID - Your Azure subscription ID. This is a GUID-formatted string (e.g. 00000000-0000-0000-0000-000000000000).
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewCertificatesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*CertificatesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &CertificatesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Description for Create or update a certificate.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-01
//   - resourceGroupName - Name of the resource group to which the resource belongs.
//   - name - Name of the certificate.
//   - certificateEnvelope - Details of certificate, if it exists already.
//   - options - CertificatesClientCreateOrUpdateOptions contains the optional parameters for the CertificatesClient.CreateOrUpdate
//     method.
func (client *CertificatesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, name string, certificateEnvelope AppCertificate, options *CertificatesClientCreateOrUpdateOptions) (CertificatesClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "CertificatesClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, name, certificateEnvelope, options)
	if err != nil {
		return CertificatesClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return CertificatesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return CertificatesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *CertificatesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, name string, certificateEnvelope AppCertificate, options *CertificatesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/certificates/{name}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, certificateEnvelope); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *CertificatesClient) createOrUpdateHandleResponse(resp *http.Response) (CertificatesClientCreateOrUpdateResponse, error) {
	result := CertificatesClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AppCertificate); err != nil {
		return CertificatesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Description for Delete a certificate.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-01
//   - resourceGroupName - Name of the resource group to which the resource belongs.
//   - name - Name of the certificate.
//   - options - CertificatesClientDeleteOptions contains the optional parameters for the CertificatesClient.Delete method.
func (client *CertificatesClient) Delete(ctx context.Context, resourceGroupName string, name string, options *CertificatesClientDeleteOptions) (CertificatesClientDeleteResponse, error) {
	var err error
	const operationName = "CertificatesClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, name, options)
	if err != nil {
		return CertificatesClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return CertificatesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return CertificatesClientDeleteResponse{}, err
	}
	return CertificatesClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *CertificatesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, name string, options *CertificatesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/certificates/{name}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Description for Get a certificate.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-01
//   - resourceGroupName - Name of the resource group to which the resource belongs.
//   - name - Name of the certificate.
//   - options - CertificatesClientGetOptions contains the optional parameters for the CertificatesClient.Get method.
func (client *CertificatesClient) Get(ctx context.Context, resourceGroupName string, name string, options *CertificatesClientGetOptions) (CertificatesClientGetResponse, error) {
	var err error
	const operationName = "CertificatesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, name, options)
	if err != nil {
		return CertificatesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return CertificatesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return CertificatesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *CertificatesClient) getCreateRequest(ctx context.Context, resourceGroupName string, name string, options *CertificatesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/certificates/{name}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *CertificatesClient) getHandleResponse(resp *http.Response) (CertificatesClientGetResponse, error) {
	result := CertificatesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AppCertificate); err != nil {
		return CertificatesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Description for Get all certificates for a subscription.
//
// Generated from API version 2024-03-01
//   - options - CertificatesClientListOptions contains the optional parameters for the CertificatesClient.NewListPager method.
func (client *CertificatesClient) NewListPager(options *CertificatesClientListOptions) *runtime.Pager[CertificatesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[CertificatesClientListResponse]{
		More: func(page CertificatesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *CertificatesClientListResponse) (CertificatesClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "CertificatesClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return CertificatesClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *CertificatesClient) listCreateRequest(ctx context.Context, options *CertificatesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Web/certificates"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	unencodedParams := []string{req.Raw().URL.RawQuery}
	if options != nil && options.Filter != nil {
		unencodedParams = append(unencodedParams, "$filter="+*options.Filter)
	}
	req.Raw().URL.RawQuery = strings.Join(unencodedParams, "&")
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *CertificatesClient) listHandleResponse(resp *http.Response) (CertificatesClientListResponse, error) {
	result := CertificatesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AppCertificateCollection); err != nil {
		return CertificatesClientListResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Description for Get all certificates in a resource group.
//
// Generated from API version 2024-03-01
//   - resourceGroupName - Name of the resource group to which the resource belongs.
//   - options - CertificatesClientListByResourceGroupOptions contains the optional parameters for the CertificatesClient.NewListByResourceGroupPager
//     method.
func (client *CertificatesClient) NewListByResourceGroupPager(resourceGroupName string, options *CertificatesClientListByResourceGroupOptions) *runtime.Pager[CertificatesClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[CertificatesClientListByResourceGroupResponse]{
		More: func(page CertificatesClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *CertificatesClientListByResourceGroupResponse) (CertificatesClientListByResourceGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "CertificatesClient.NewListByResourceGroupPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			}, nil)
			if err != nil {
				return CertificatesClientListByResourceGroupResponse{}, err
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *CertificatesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *CertificatesClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/certificates"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *CertificatesClient) listByResourceGroupHandleResponse(resp *http.Response) (CertificatesClientListByResourceGroupResponse, error) {
	result := CertificatesClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AppCertificateCollection); err != nil {
		return CertificatesClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// Update - Description for Create or update a certificate.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-01
//   - resourceGroupName - Name of the resource group to which the resource belongs.
//   - name - Name of the certificate.
//   - certificateEnvelope - Details of certificate, if it exists already.
//   - options - CertificatesClientUpdateOptions contains the optional parameters for the CertificatesClient.Update method.
func (client *CertificatesClient) Update(ctx context.Context, resourceGroupName string, name string, certificateEnvelope AppCertificatePatchResource, options *CertificatesClientUpdateOptions) (CertificatesClientUpdateResponse, error) {
	var err error
	const operationName = "CertificatesClient.Update"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, name, certificateEnvelope, options)
	if err != nil {
		return CertificatesClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return CertificatesClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return CertificatesClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *CertificatesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, name string, certificateEnvelope AppCertificatePatchResource, options *CertificatesClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/certificates/{name}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, certificateEnvelope); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *CertificatesClient) updateHandleResponse(resp *http.Response) (CertificatesClientUpdateResponse, error) {
	result := CertificatesClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AppCertificate); err != nil {
		return CertificatesClientUpdateResponse{}, err
	}
	return result, nil
}
