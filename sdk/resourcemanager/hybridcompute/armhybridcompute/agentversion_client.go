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
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// AgentVersionClient contains the methods for the AgentVersion group.
// Don't use this type directly, use NewAgentVersionClient() instead.
type AgentVersionClient struct {
	internal *arm.Client
}

// NewAgentVersionClient creates a new instance of AgentVersionClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewAgentVersionClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*AgentVersionClient, error) {
	cl, err := arm.NewClient(moduleName+".AgentVersionClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &AgentVersionClient{
		internal: cl,
	}
	return client, nil
}

// Get - Gets an Agent Version along with the download link currently present.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-25-preview
//   - osType - Defines the os type
//   - version - Defines the agent version. To get latest, use latest or else a specific agent version.
//   - options - AgentVersionClientGetOptions contains the optional parameters for the AgentVersionClient.Get method.
func (client *AgentVersionClient) Get(ctx context.Context, osType string, version string, options *AgentVersionClientGetOptions) (AgentVersionClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, osType, version, options)
	if err != nil {
		return AgentVersionClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AgentVersionClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AgentVersionClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *AgentVersionClient) getCreateRequest(ctx context.Context, osType string, version string, options *AgentVersionClientGetOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.HybridCompute/osType/{osType}/agentVersions/{version}"
	if osType == "" {
		return nil, errors.New("parameter osType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{osType}", url.PathEscape(osType))
	if version == "" {
		return nil, errors.New("parameter version cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{version}", url.PathEscape(version))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-25-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AgentVersionClient) getHandleResponse(resp *http.Response) (AgentVersionClientGetResponse, error) {
	result := AgentVersionClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AgentVersion); err != nil {
		return AgentVersionClientGetResponse{}, err
	}
	return result, nil
}

// List - Gets all Agent Versions along with the download link currently present.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-25-preview
//   - osType - Defines the os type.
//   - options - AgentVersionClientListOptions contains the optional parameters for the AgentVersionClient.List method.
func (client *AgentVersionClient) List(ctx context.Context, osType string, options *AgentVersionClientListOptions) (AgentVersionClientListResponse, error) {
	req, err := client.listCreateRequest(ctx, osType, options)
	if err != nil {
		return AgentVersionClientListResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AgentVersionClientListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AgentVersionClientListResponse{}, runtime.NewResponseError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *AgentVersionClient) listCreateRequest(ctx context.Context, osType string, options *AgentVersionClientListOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.HybridCompute/osType/{osType}/agentVersions"
	if osType == "" {
		return nil, errors.New("parameter osType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{osType}", url.PathEscape(osType))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-25-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *AgentVersionClient) listHandleResponse(resp *http.Response) (AgentVersionClientListResponse, error) {
	result := AgentVersionClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AgentVersionsList); err != nil {
		return AgentVersionClientListResponse{}, err
	}
	return result, nil
}
