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

// TopLevelDomainsClient contains the methods for the TopLevelDomains group.
// Don't use this type directly, use NewTopLevelDomainsClient() instead.
type TopLevelDomainsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewTopLevelDomainsClient creates a new instance of TopLevelDomainsClient with the specified values.
//   - subscriptionID - Your Azure subscription ID. This is a GUID-formatted string (e.g. 00000000-0000-0000-0000-000000000000).
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewTopLevelDomainsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*TopLevelDomainsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &TopLevelDomainsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Description for Get details of a top-level domain.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-12-01
//   - name - Name of the top-level domain.
//   - options - TopLevelDomainsClientGetOptions contains the optional parameters for the TopLevelDomainsClient.Get method.
func (client *TopLevelDomainsClient) Get(ctx context.Context, name string, options *TopLevelDomainsClientGetOptions) (TopLevelDomainsClientGetResponse, error) {
	var err error
	const operationName = "TopLevelDomainsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, name, options)
	if err != nil {
		return TopLevelDomainsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TopLevelDomainsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return TopLevelDomainsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *TopLevelDomainsClient) getCreateRequest(ctx context.Context, name string, options *TopLevelDomainsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.DomainRegistration/topLevelDomains/{name}"
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
	reqQP.Set("api-version", "2023-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *TopLevelDomainsClient) getHandleResponse(resp *http.Response) (TopLevelDomainsClientGetResponse, error) {
	result := TopLevelDomainsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TopLevelDomain); err != nil {
		return TopLevelDomainsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Description for Get all top-level domains supported for registration.
//
// Generated from API version 2023-12-01
//   - options - TopLevelDomainsClientListOptions contains the optional parameters for the TopLevelDomainsClient.NewListPager
//     method.
func (client *TopLevelDomainsClient) NewListPager(options *TopLevelDomainsClientListOptions) *runtime.Pager[TopLevelDomainsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[TopLevelDomainsClientListResponse]{
		More: func(page TopLevelDomainsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *TopLevelDomainsClientListResponse) (TopLevelDomainsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "TopLevelDomainsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return TopLevelDomainsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *TopLevelDomainsClient) listCreateRequest(ctx context.Context, options *TopLevelDomainsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.DomainRegistration/topLevelDomains"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *TopLevelDomainsClient) listHandleResponse(resp *http.Response) (TopLevelDomainsClientListResponse, error) {
	result := TopLevelDomainsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TopLevelDomainCollection); err != nil {
		return TopLevelDomainsClientListResponse{}, err
	}
	return result, nil
}

// NewListAgreementsPager - Description for Gets all legal agreements that user needs to accept before purchasing a domain.
//
// Generated from API version 2023-12-01
//   - name - Name of the top-level domain.
//   - agreementOption - Domain agreement options.
//   - options - TopLevelDomainsClientListAgreementsOptions contains the optional parameters for the TopLevelDomainsClient.NewListAgreementsPager
//     method.
func (client *TopLevelDomainsClient) NewListAgreementsPager(name string, agreementOption TopLevelDomainAgreementOption, options *TopLevelDomainsClientListAgreementsOptions) *runtime.Pager[TopLevelDomainsClientListAgreementsResponse] {
	return runtime.NewPager(runtime.PagingHandler[TopLevelDomainsClientListAgreementsResponse]{
		More: func(page TopLevelDomainsClientListAgreementsResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *TopLevelDomainsClientListAgreementsResponse) (TopLevelDomainsClientListAgreementsResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "TopLevelDomainsClient.NewListAgreementsPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listAgreementsCreateRequest(ctx, name, agreementOption, options)
			}, nil)
			if err != nil {
				return TopLevelDomainsClientListAgreementsResponse{}, err
			}
			return client.listAgreementsHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listAgreementsCreateRequest creates the ListAgreements request.
func (client *TopLevelDomainsClient) listAgreementsCreateRequest(ctx context.Context, name string, agreementOption TopLevelDomainAgreementOption, options *TopLevelDomainsClientListAgreementsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.DomainRegistration/topLevelDomains/{name}/listAgreements"
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, agreementOption); err != nil {
		return nil, err
	}
	return req, nil
}

// listAgreementsHandleResponse handles the ListAgreements response.
func (client *TopLevelDomainsClient) listAgreementsHandleResponse(resp *http.Response) (TopLevelDomainsClientListAgreementsResponse, error) {
	result := TopLevelDomainsClientListAgreementsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TldLegalAgreementCollection); err != nil {
		return TopLevelDomainsClientListAgreementsResponse{}, err
	}
	return result, nil
}
