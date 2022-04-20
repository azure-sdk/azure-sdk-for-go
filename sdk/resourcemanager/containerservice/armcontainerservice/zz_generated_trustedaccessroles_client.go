//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcontainerservice

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

// TrustedAccessRolesClient contains the methods for the TrustedAccessRoles group.
// Don't use this type directly, use NewTrustedAccessRolesClient() instead.
type TrustedAccessRolesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewTrustedAccessRolesClient creates a new instance of TrustedAccessRolesClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewTrustedAccessRolesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*TrustedAccessRolesClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublicCloud.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &TrustedAccessRolesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// List - List supported trusted access roles.
// If the operation fails it returns an *azcore.ResponseError type.
// location - The name of Azure region.
// options - TrustedAccessRolesClientListOptions contains the optional parameters for the TrustedAccessRolesClient.List method.
func (client *TrustedAccessRolesClient) List(ctx context.Context, location string, options *TrustedAccessRolesClientListOptions) (TrustedAccessRolesClientListResponse, error) {
	req, err := client.listCreateRequest(ctx, location, options)
	if err != nil {
		return TrustedAccessRolesClientListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TrustedAccessRolesClientListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return TrustedAccessRolesClientListResponse{}, runtime.NewResponseError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *TrustedAccessRolesClient) listCreateRequest(ctx context.Context, location string, options *TrustedAccessRolesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ContainerService/locations/{location}/trustedAccessRoles"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-04-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *TrustedAccessRolesClient) listHandleResponse(resp *http.Response) (TrustedAccessRolesClientListResponse, error) {
	result := TrustedAccessRolesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TrustedAccessRoleListResult); err != nil {
		return TrustedAccessRolesClientListResponse{}, err
	}
	return result, nil
}
