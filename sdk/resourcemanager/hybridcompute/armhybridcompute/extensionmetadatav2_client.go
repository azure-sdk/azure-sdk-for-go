// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhybridcompute

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

// ExtensionMetadataV2Client contains the methods for the ExtensionMetadataV2 group.
// Don't use this type directly, use NewExtensionMetadataV2Client() instead.
type ExtensionMetadataV2Client struct {
	internal *arm.Client
}

// NewExtensionMetadataV2Client creates a new instance of ExtensionMetadataV2Client with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewExtensionMetadataV2Client(credential azcore.TokenCredential, options *arm.ClientOptions) (*ExtensionMetadataV2Client, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ExtensionMetadataV2Client{
		internal: cl,
	}
	return client, nil
}

// Get - Gets an Extension Metadata based on location, publisher, extensionType and version
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-02-19-preview
//   - location - The name of Azure region.
//   - publisher - The publisher of the Extension being received.
//   - extensionType - The extensionType of the Extension being received.
//   - version - The version of the Extension being received.
//   - options - ExtensionMetadataV2ClientGetOptions contains the optional parameters for the ExtensionMetadataV2Client.Get method.
func (client *ExtensionMetadataV2Client) Get(ctx context.Context, location string, publisher string, extensionType string, version string, options *ExtensionMetadataV2ClientGetOptions) (ExtensionMetadataV2ClientGetResponse, error) {
	var err error
	const operationName = "ExtensionMetadataV2Client.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, location, publisher, extensionType, version, options)
	if err != nil {
		return ExtensionMetadataV2ClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ExtensionMetadataV2ClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ExtensionMetadataV2ClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ExtensionMetadataV2Client) getCreateRequest(ctx context.Context, location string, publisher string, extensionType string, version string, _ *ExtensionMetadataV2ClientGetOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.HybridCompute/locations/{location}/publishers/{publisher}/extensionTypes/{extensionType}/versions/{version}"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if publisher == "" {
		return nil, errors.New("parameter publisher cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{publisher}", url.PathEscape(publisher))
	if extensionType == "" {
		return nil, errors.New("parameter extensionType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{extensionType}", url.PathEscape(extensionType))
	if version == "" {
		return nil, errors.New("parameter version cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{version}", url.PathEscape(version))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-02-19-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ExtensionMetadataV2Client) getHandleResponse(resp *http.Response) (ExtensionMetadataV2ClientGetResponse, error) {
	result := ExtensionMetadataV2ClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExtensionValueV2); err != nil {
		return ExtensionMetadataV2ClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets all Extension versions based on location, publisher, extensionType
//
// Generated from API version 2025-02-19-preview
//   - location - The name of Azure region.
//   - publisher - The publisher of the Extension being received.
//   - extensionType - The extensionType of the Extension being received.
//   - options - ExtensionMetadataV2ClientListOptions contains the optional parameters for the ExtensionMetadataV2Client.NewListPager
//     method.
func (client *ExtensionMetadataV2Client) NewListPager(location string, publisher string, extensionType string, options *ExtensionMetadataV2ClientListOptions) *runtime.Pager[ExtensionMetadataV2ClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ExtensionMetadataV2ClientListResponse]{
		More: func(page ExtensionMetadataV2ClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ExtensionMetadataV2ClientListResponse) (ExtensionMetadataV2ClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ExtensionMetadataV2Client.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, location, publisher, extensionType, options)
			}, nil)
			if err != nil {
				return ExtensionMetadataV2ClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *ExtensionMetadataV2Client) listCreateRequest(ctx context.Context, location string, publisher string, extensionType string, _ *ExtensionMetadataV2ClientListOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.HybridCompute/locations/{location}/publishers/{publisher}/extensionTypes/{extensionType}/versions"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if publisher == "" {
		return nil, errors.New("parameter publisher cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{publisher}", url.PathEscape(publisher))
	if extensionType == "" {
		return nil, errors.New("parameter extensionType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{extensionType}", url.PathEscape(extensionType))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-02-19-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ExtensionMetadataV2Client) listHandleResponse(resp *http.Response) (ExtensionMetadataV2ClientListResponse, error) {
	result := ExtensionMetadataV2ClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExtensionValueListResultV2); err != nil {
		return ExtensionMetadataV2ClientListResponse{}, err
	}
	return result, nil
}
