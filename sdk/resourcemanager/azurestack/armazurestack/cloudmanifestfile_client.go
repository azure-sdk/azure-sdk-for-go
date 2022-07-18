//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armazurestack

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

// CloudManifestFileClient contains the methods for the CloudManifestFile group.
// Don't use this type directly, use NewCloudManifestFileClient() instead.
type CloudManifestFileClient struct {
	host string
	pl   runtime.Pipeline
}

// NewCloudManifestFileClient creates a new instance of CloudManifestFileClient with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewCloudManifestFileClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*CloudManifestFileClient, error) {
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
	client := &CloudManifestFileClient{
		host: ep,
		pl:   pl,
	}
	return client, nil
}

// Get - Returns a cloud specific manifest JSON file.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-06-01
// verificationVersion - Signing verification key version.
// options - CloudManifestFileClientGetOptions contains the optional parameters for the CloudManifestFileClient.Get method.
func (client *CloudManifestFileClient) Get(ctx context.Context, verificationVersion string, options *CloudManifestFileClientGetOptions) (CloudManifestFileClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, verificationVersion, options)
	if err != nil {
		return CloudManifestFileClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CloudManifestFileClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return CloudManifestFileClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *CloudManifestFileClient) getCreateRequest(ctx context.Context, verificationVersion string, options *CloudManifestFileClientGetOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.AzureStack/cloudManifestFiles/{verificationVersion}"
	if verificationVersion == "" {
		return nil, errors.New("parameter verificationVersion cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{verificationVersion}", url.PathEscape(verificationVersion))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.VersionCreationDate != nil {
		reqQP.Set("versionCreationDate", *options.VersionCreationDate)
	}
	reqQP.Set("api-version", "2022-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *CloudManifestFileClient) getHandleResponse(resp *http.Response) (CloudManifestFileClientGetResponse, error) {
	result := CloudManifestFileClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CloudManifestFileResponse); err != nil {
		return CloudManifestFileClientGetResponse{}, err
	}
	return result, nil
}

// List - Returns a cloud specific manifest JSON file with latest version.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-06-01
// options - CloudManifestFileClientListOptions contains the optional parameters for the CloudManifestFileClient.List method.
func (client *CloudManifestFileClient) List(ctx context.Context, options *CloudManifestFileClientListOptions) (CloudManifestFileClientListResponse, error) {
	req, err := client.listCreateRequest(ctx, options)
	if err != nil {
		return CloudManifestFileClientListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CloudManifestFileClientListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return CloudManifestFileClientListResponse{}, runtime.NewResponseError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *CloudManifestFileClient) listCreateRequest(ctx context.Context, options *CloudManifestFileClientListOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.AzureStack/cloudManifestFiles"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *CloudManifestFileClient) listHandleResponse(resp *http.Response) (CloudManifestFileClientListResponse, error) {
	result := CloudManifestFileClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CloudManifestFileResponse); err != nil {
		return CloudManifestFileClientListResponse{}, err
	}
	return result, nil
}
