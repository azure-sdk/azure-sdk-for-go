//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsupport

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

// ServicesClient contains the methods for the Services group.
// Don't use this type directly, use NewServicesClient() instead.
type ServicesClient struct {
	internal *arm.Client
}

// NewServicesClient creates a new instance of ServicesClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewServicesClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*ServicesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ServicesClient{
		internal: cl,
	}
	return client, nil
}

// Get - Gets a specific Azure service for support ticket creation.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01-preview
//   - serviceName - Name of the Azure service for which the problem classifications need to be retrieved.
//   - options - ServicesClientGetOptions contains the optional parameters for the ServicesClient.Get method.
func (client *ServicesClient) Get(ctx context.Context, serviceName string, options *ServicesClientGetOptions) (ServicesClientGetResponse, error) {
	var err error
	const operationName = "ServicesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, serviceName, options)
	if err != nil {
		return ServicesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ServicesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ServicesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ServicesClient) getCreateRequest(ctx context.Context, serviceName string, options *ServicesClientGetOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Support/services/{serviceName}"
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ServicesClient) getHandleResponse(resp *http.Response) (ServicesClientGetResponse, error) {
	result := ServicesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Service); err != nil {
		return ServicesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists all the Azure services available for support ticket creation. For Technical issues, select the Service
// Id that maps to the Azure service/product as displayed in the Services drop-down list on
// the Azure portal's New support request [https://portal.azure.com/#blade/MicrosoftAzureSupport/HelpAndSupportBlade/overview]
// page. Always use the service and its corresponding problem classification(s)
// obtained programmatically for support ticket creation. This practice ensures that you always have the most recent set of
// service and problem classification Ids.
//
// Generated from API version 2023-06-01-preview
//   - options - ServicesClientListOptions contains the optional parameters for the ServicesClient.NewListPager method.
func (client *ServicesClient) NewListPager(options *ServicesClientListOptions) *runtime.Pager[ServicesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ServicesClientListResponse]{
		More: func(page ServicesClientListResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *ServicesClientListResponse) (ServicesClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ServicesClient.NewListPager")
			req, err := client.listCreateRequest(ctx, options)
			if err != nil {
				return ServicesClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ServicesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ServicesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *ServicesClient) listCreateRequest(ctx context.Context, options *ServicesClientListOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Support/services"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ServicesClient) listHandleResponse(resp *http.Response) (ServicesClientListResponse, error) {
	result := ServicesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServicesListResult); err != nil {
		return ServicesClientListResponse{}, err
	}
	return result, nil
}
