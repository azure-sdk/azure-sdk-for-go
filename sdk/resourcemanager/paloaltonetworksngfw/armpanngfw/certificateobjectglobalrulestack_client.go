//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpanngfw

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

// CertificateObjectGlobalRulestackClient contains the methods for the CertificateObjectGlobalRulestack group.
// Don't use this type directly, use NewCertificateObjectGlobalRulestackClient() instead.
type CertificateObjectGlobalRulestackClient struct {
	internal *arm.Client
}

// NewCertificateObjectGlobalRulestackClient creates a new instance of CertificateObjectGlobalRulestackClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewCertificateObjectGlobalRulestackClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*CertificateObjectGlobalRulestackClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &CertificateObjectGlobalRulestackClient{
		internal: cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create a CertificateObjectGlobalRulestackResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-19-preview
//   - globalRulestackName - GlobalRulestack resource name
//   - name - certificate name
//   - resource - Resource create parameters.
//   - options - CertificateObjectGlobalRulestackClientBeginCreateOrUpdateOptions contains the optional parameters for the CertificateObjectGlobalRulestackClient.BeginCreateOrUpdate
//     method.
func (client *CertificateObjectGlobalRulestackClient) BeginCreateOrUpdate(ctx context.Context, globalRulestackName string, name string, resource CertificateObjectGlobalRulestackResource, options *CertificateObjectGlobalRulestackClientBeginCreateOrUpdateOptions) (*runtime.Poller[CertificateObjectGlobalRulestackClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, globalRulestackName, name, resource, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[CertificateObjectGlobalRulestackClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[CertificateObjectGlobalRulestackClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Create a CertificateObjectGlobalRulestackResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-19-preview
func (client *CertificateObjectGlobalRulestackClient) createOrUpdate(ctx context.Context, globalRulestackName string, name string, resource CertificateObjectGlobalRulestackResource, options *CertificateObjectGlobalRulestackClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "CertificateObjectGlobalRulestackClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, globalRulestackName, name, resource, options)
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
func (client *CertificateObjectGlobalRulestackClient) createOrUpdateCreateRequest(ctx context.Context, globalRulestackName string, name string, resource CertificateObjectGlobalRulestackResource, options *CertificateObjectGlobalRulestackClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/providers/PaloAltoNetworks.Cloudngfw/globalRulestacks/{globalRulestackName}/certificates/{name}"
	if globalRulestackName == "" {
		return nil, errors.New("parameter globalRulestackName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{globalRulestackName}", url.PathEscape(globalRulestackName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-01-19-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, resource); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Delete a CertificateObjectGlobalRulestackResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-19-preview
//   - globalRulestackName - GlobalRulestack resource name
//   - name - certificate name
//   - options - CertificateObjectGlobalRulestackClientBeginDeleteOptions contains the optional parameters for the CertificateObjectGlobalRulestackClient.BeginDelete
//     method.
func (client *CertificateObjectGlobalRulestackClient) BeginDelete(ctx context.Context, globalRulestackName string, name string, options *CertificateObjectGlobalRulestackClientBeginDeleteOptions) (*runtime.Poller[CertificateObjectGlobalRulestackClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, globalRulestackName, name, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[CertificateObjectGlobalRulestackClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[CertificateObjectGlobalRulestackClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Delete a CertificateObjectGlobalRulestackResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-19-preview
func (client *CertificateObjectGlobalRulestackClient) deleteOperation(ctx context.Context, globalRulestackName string, name string, options *CertificateObjectGlobalRulestackClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "CertificateObjectGlobalRulestackClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, globalRulestackName, name, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *CertificateObjectGlobalRulestackClient) deleteCreateRequest(ctx context.Context, globalRulestackName string, name string, options *CertificateObjectGlobalRulestackClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/providers/PaloAltoNetworks.Cloudngfw/globalRulestacks/{globalRulestackName}/certificates/{name}"
	if globalRulestackName == "" {
		return nil, errors.New("parameter globalRulestackName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{globalRulestackName}", url.PathEscape(globalRulestackName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-01-19-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get a CertificateObjectGlobalRulestackResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-19-preview
//   - globalRulestackName - GlobalRulestack resource name
//   - name - certificate name
//   - options - CertificateObjectGlobalRulestackClientGetOptions contains the optional parameters for the CertificateObjectGlobalRulestackClient.Get
//     method.
func (client *CertificateObjectGlobalRulestackClient) Get(ctx context.Context, globalRulestackName string, name string, options *CertificateObjectGlobalRulestackClientGetOptions) (CertificateObjectGlobalRulestackClientGetResponse, error) {
	var err error
	const operationName = "CertificateObjectGlobalRulestackClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, globalRulestackName, name, options)
	if err != nil {
		return CertificateObjectGlobalRulestackClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return CertificateObjectGlobalRulestackClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return CertificateObjectGlobalRulestackClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *CertificateObjectGlobalRulestackClient) getCreateRequest(ctx context.Context, globalRulestackName string, name string, options *CertificateObjectGlobalRulestackClientGetOptions) (*policy.Request, error) {
	urlPath := "/providers/PaloAltoNetworks.Cloudngfw/globalRulestacks/{globalRulestackName}/certificates/{name}"
	if globalRulestackName == "" {
		return nil, errors.New("parameter globalRulestackName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{globalRulestackName}", url.PathEscape(globalRulestackName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-01-19-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *CertificateObjectGlobalRulestackClient) getHandleResponse(resp *http.Response) (CertificateObjectGlobalRulestackClientGetResponse, error) {
	result := CertificateObjectGlobalRulestackClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CertificateObjectGlobalRulestackResource); err != nil {
		return CertificateObjectGlobalRulestackClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List CertificateObjectGlobalRulestackResource resources by Tenant
//
// Generated from API version 2024-01-19-preview
//   - globalRulestackName - GlobalRulestack resource name
//   - options - CertificateObjectGlobalRulestackClientListOptions contains the optional parameters for the CertificateObjectGlobalRulestackClient.NewListPager
//     method.
func (client *CertificateObjectGlobalRulestackClient) NewListPager(globalRulestackName string, options *CertificateObjectGlobalRulestackClientListOptions) *runtime.Pager[CertificateObjectGlobalRulestackClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[CertificateObjectGlobalRulestackClientListResponse]{
		More: func(page CertificateObjectGlobalRulestackClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *CertificateObjectGlobalRulestackClientListResponse) (CertificateObjectGlobalRulestackClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "CertificateObjectGlobalRulestackClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, globalRulestackName, options)
			}, nil)
			if err != nil {
				return CertificateObjectGlobalRulestackClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *CertificateObjectGlobalRulestackClient) listCreateRequest(ctx context.Context, globalRulestackName string, options *CertificateObjectGlobalRulestackClientListOptions) (*policy.Request, error) {
	urlPath := "/providers/PaloAltoNetworks.Cloudngfw/globalRulestacks/{globalRulestackName}/certificates"
	if globalRulestackName == "" {
		return nil, errors.New("parameter globalRulestackName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{globalRulestackName}", url.PathEscape(globalRulestackName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-01-19-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *CertificateObjectGlobalRulestackClient) listHandleResponse(resp *http.Response) (CertificateObjectGlobalRulestackClientListResponse, error) {
	result := CertificateObjectGlobalRulestackClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CertificateObjectGlobalRulestackResourceListResult); err != nil {
		return CertificateObjectGlobalRulestackClientListResponse{}, err
	}
	return result, nil
}
