//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armhybridcompute

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
	"strings"
)

// ExtensionMetadataClient contains the methods for the ExtensionMetadata group.
// Don't use this type directly, use NewExtensionMetadataClient() instead.
type ExtensionMetadataClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewExtensionMetadataClient creates a new instance of ExtensionMetadataClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewExtensionMetadataClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ExtensionMetadataClient, error) {
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
	client := &ExtensionMetadataClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Get - Gets an Extension Metadata based on location, publisher, extensionType and version
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-03-preview
//   - location - The location of the Extension being received.
//   - publisher - The publisher of the Extension being received.
//   - extensionType - The extensionType of the Extension being received.
//   - version - The version of the Extension being received.
//   - options - ExtensionMetadataClientGetOptions contains the optional parameters for the ExtensionMetadataClient.Get method.
func (client *ExtensionMetadataClient) Get(ctx context.Context, location string, publisher string, extensionType string, version string, options *ExtensionMetadataClientGetOptions) (ExtensionMetadataClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, location, publisher, extensionType, version, options)
	if err != nil {
		return ExtensionMetadataClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ExtensionMetadataClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ExtensionMetadataClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ExtensionMetadataClient) getCreateRequest(ctx context.Context, location string, publisher string, extensionType string, version string, options *ExtensionMetadataClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.HybridCompute/locations/{location}/publishers/{publisher}/extensionTypes/{extensionType}/versions/{version}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-03-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ExtensionMetadataClient) getHandleResponse(resp *http.Response) (ExtensionMetadataClientGetResponse, error) {
	result := ExtensionMetadataClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExtensionValue); err != nil {
		return ExtensionMetadataClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets all Extension versions based on location, publisher, extensionType
//
// Generated from API version 2023-03-03-preview
//   - location - The location of the Extension being received.
//   - publisher - The publisher of the Extension being received.
//   - extensionType - The extensionType of the Extension being received.
//   - options - ExtensionMetadataClientListOptions contains the optional parameters for the ExtensionMetadataClient.NewListPager
//     method.
func (client *ExtensionMetadataClient) NewListPager(location string, publisher string, extensionType string, options *ExtensionMetadataClientListOptions) *runtime.Pager[ExtensionMetadataClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ExtensionMetadataClientListResponse]{
		More: func(page ExtensionMetadataClientListResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *ExtensionMetadataClientListResponse) (ExtensionMetadataClientListResponse, error) {
			req, err := client.listCreateRequest(ctx, location, publisher, extensionType, options)
			if err != nil {
				return ExtensionMetadataClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ExtensionMetadataClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ExtensionMetadataClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *ExtensionMetadataClient) listCreateRequest(ctx context.Context, location string, publisher string, extensionType string, options *ExtensionMetadataClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.HybridCompute/locations/{location}/publishers/{publisher}/extensionTypes/{extensionType}/versions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-03-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ExtensionMetadataClient) listHandleResponse(resp *http.Response) (ExtensionMetadataClientListResponse, error) {
	result := ExtensionMetadataClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExtensionValueListResult); err != nil {
		return ExtensionMetadataClientListResponse{}, err
	}
	return result, nil
}
