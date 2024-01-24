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

// CertificateObjectLocalRulestackClient contains the methods for the CertificateObjectLocalRulestack group.
// Don't use this type directly, use NewCertificateObjectLocalRulestackClient() instead.
type CertificateObjectLocalRulestackClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewCertificateObjectLocalRulestackClient creates a new instance of CertificateObjectLocalRulestackClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewCertificateObjectLocalRulestackClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*CertificateObjectLocalRulestackClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &CertificateObjectLocalRulestackClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create a CertificateObjectLocalRulestackResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-19-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - localRulestackName - LocalRulestack resource name
//   - name - certificate name
//   - resource - Resource create parameters.
//   - options - CertificateObjectLocalRulestackClientBeginCreateOrUpdateOptions contains the optional parameters for the CertificateObjectLocalRulestackClient.BeginCreateOrUpdate
//     method.
func (client *CertificateObjectLocalRulestackClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, localRulestackName string, name string, resource CertificateObjectLocalRulestackResource, options *CertificateObjectLocalRulestackClientBeginCreateOrUpdateOptions) (*runtime.Poller[CertificateObjectLocalRulestackClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, localRulestackName, name, resource, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[CertificateObjectLocalRulestackClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[CertificateObjectLocalRulestackClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Create a CertificateObjectLocalRulestackResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-19-preview
func (client *CertificateObjectLocalRulestackClient) createOrUpdate(ctx context.Context, resourceGroupName string, localRulestackName string, name string, resource CertificateObjectLocalRulestackResource, options *CertificateObjectLocalRulestackClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "CertificateObjectLocalRulestackClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, localRulestackName, name, resource, options)
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
func (client *CertificateObjectLocalRulestackClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, localRulestackName string, name string, resource CertificateObjectLocalRulestackResource, options *CertificateObjectLocalRulestackClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/PaloAltoNetworks.Cloudngfw/localRulestacks/{localRulestackName}/certificates/{name}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if localRulestackName == "" {
		return nil, errors.New("parameter localRulestackName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{localRulestackName}", url.PathEscape(localRulestackName))
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

// BeginDelete - Delete a CertificateObjectLocalRulestackResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-19-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - localRulestackName - LocalRulestack resource name
//   - name - certificate name
//   - options - CertificateObjectLocalRulestackClientBeginDeleteOptions contains the optional parameters for the CertificateObjectLocalRulestackClient.BeginDelete
//     method.
func (client *CertificateObjectLocalRulestackClient) BeginDelete(ctx context.Context, resourceGroupName string, localRulestackName string, name string, options *CertificateObjectLocalRulestackClientBeginDeleteOptions) (*runtime.Poller[CertificateObjectLocalRulestackClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, localRulestackName, name, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[CertificateObjectLocalRulestackClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[CertificateObjectLocalRulestackClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Delete a CertificateObjectLocalRulestackResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-19-preview
func (client *CertificateObjectLocalRulestackClient) deleteOperation(ctx context.Context, resourceGroupName string, localRulestackName string, name string, options *CertificateObjectLocalRulestackClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "CertificateObjectLocalRulestackClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, localRulestackName, name, options)
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
func (client *CertificateObjectLocalRulestackClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, localRulestackName string, name string, options *CertificateObjectLocalRulestackClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/PaloAltoNetworks.Cloudngfw/localRulestacks/{localRulestackName}/certificates/{name}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if localRulestackName == "" {
		return nil, errors.New("parameter localRulestackName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{localRulestackName}", url.PathEscape(localRulestackName))
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

// Get - Get a CertificateObjectLocalRulestackResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-19-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - localRulestackName - LocalRulestack resource name
//   - name - certificate name
//   - options - CertificateObjectLocalRulestackClientGetOptions contains the optional parameters for the CertificateObjectLocalRulestackClient.Get
//     method.
func (client *CertificateObjectLocalRulestackClient) Get(ctx context.Context, resourceGroupName string, localRulestackName string, name string, options *CertificateObjectLocalRulestackClientGetOptions) (CertificateObjectLocalRulestackClientGetResponse, error) {
	var err error
	const operationName = "CertificateObjectLocalRulestackClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, localRulestackName, name, options)
	if err != nil {
		return CertificateObjectLocalRulestackClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return CertificateObjectLocalRulestackClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return CertificateObjectLocalRulestackClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *CertificateObjectLocalRulestackClient) getCreateRequest(ctx context.Context, resourceGroupName string, localRulestackName string, name string, options *CertificateObjectLocalRulestackClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/PaloAltoNetworks.Cloudngfw/localRulestacks/{localRulestackName}/certificates/{name}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if localRulestackName == "" {
		return nil, errors.New("parameter localRulestackName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{localRulestackName}", url.PathEscape(localRulestackName))
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
func (client *CertificateObjectLocalRulestackClient) getHandleResponse(resp *http.Response) (CertificateObjectLocalRulestackClientGetResponse, error) {
	result := CertificateObjectLocalRulestackClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CertificateObjectLocalRulestackResource); err != nil {
		return CertificateObjectLocalRulestackClientGetResponse{}, err
	}
	return result, nil
}

// NewListByLocalRulestacksPager - List CertificateObjectLocalRulestackResource resources by LocalRulestacks
//
// Generated from API version 2024-01-19-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - localRulestackName - LocalRulestack resource name
//   - options - CertificateObjectLocalRulestackClientListByLocalRulestacksOptions contains the optional parameters for the CertificateObjectLocalRulestackClient.NewListByLocalRulestacksPager
//     method.
func (client *CertificateObjectLocalRulestackClient) NewListByLocalRulestacksPager(resourceGroupName string, localRulestackName string, options *CertificateObjectLocalRulestackClientListByLocalRulestacksOptions) *runtime.Pager[CertificateObjectLocalRulestackClientListByLocalRulestacksResponse] {
	return runtime.NewPager(runtime.PagingHandler[CertificateObjectLocalRulestackClientListByLocalRulestacksResponse]{
		More: func(page CertificateObjectLocalRulestackClientListByLocalRulestacksResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *CertificateObjectLocalRulestackClientListByLocalRulestacksResponse) (CertificateObjectLocalRulestackClientListByLocalRulestacksResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "CertificateObjectLocalRulestackClient.NewListByLocalRulestacksPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByLocalRulestacksCreateRequest(ctx, resourceGroupName, localRulestackName, options)
			}, nil)
			if err != nil {
				return CertificateObjectLocalRulestackClientListByLocalRulestacksResponse{}, err
			}
			return client.listByLocalRulestacksHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByLocalRulestacksCreateRequest creates the ListByLocalRulestacks request.
func (client *CertificateObjectLocalRulestackClient) listByLocalRulestacksCreateRequest(ctx context.Context, resourceGroupName string, localRulestackName string, options *CertificateObjectLocalRulestackClientListByLocalRulestacksOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/PaloAltoNetworks.Cloudngfw/localRulestacks/{localRulestackName}/certificates"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if localRulestackName == "" {
		return nil, errors.New("parameter localRulestackName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{localRulestackName}", url.PathEscape(localRulestackName))
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

// listByLocalRulestacksHandleResponse handles the ListByLocalRulestacks response.
func (client *CertificateObjectLocalRulestackClient) listByLocalRulestacksHandleResponse(resp *http.Response) (CertificateObjectLocalRulestackClientListByLocalRulestacksResponse, error) {
	result := CertificateObjectLocalRulestackClientListByLocalRulestacksResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CertificateObjectLocalRulestackResourceListResult); err != nil {
		return CertificateObjectLocalRulestackClientListByLocalRulestacksResponse{}, err
	}
	return result, nil
}
