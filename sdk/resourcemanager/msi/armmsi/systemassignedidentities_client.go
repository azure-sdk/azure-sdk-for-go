//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmsi

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strings"
)

// SystemAssignedIdentitiesClient contains the methods for the SystemAssignedIdentities group.
// Don't use this type directly, use NewSystemAssignedIdentitiesClient() instead.
type SystemAssignedIdentitiesClient struct {
	internal *arm.Client
}

// NewSystemAssignedIdentitiesClient creates a new instance of SystemAssignedIdentitiesClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewSystemAssignedIdentitiesClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*SystemAssignedIdentitiesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &SystemAssignedIdentitiesClient{
		internal: cl,
	}
	return client, nil
}

// GetByScope - Gets the systemAssignedIdentity available under the specified RP scope.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-11-30
//   - scope - The resource provider scope of the resource. Parent resource being extended by Managed Identities.
//   - options - SystemAssignedIdentitiesClientGetByScopeOptions contains the optional parameters for the SystemAssignedIdentitiesClient.GetByScope
//     method.
func (client *SystemAssignedIdentitiesClient) GetByScope(ctx context.Context, scope string, options *SystemAssignedIdentitiesClientGetByScopeOptions) (SystemAssignedIdentitiesClientGetByScopeResponse, error) {
	var err error
	const operationName = "SystemAssignedIdentitiesClient.GetByScope"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getByScopeCreateRequest(ctx, scope, options)
	if err != nil {
		return SystemAssignedIdentitiesClientGetByScopeResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SystemAssignedIdentitiesClientGetByScopeResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SystemAssignedIdentitiesClientGetByScopeResponse{}, err
	}
	resp, err := client.getByScopeHandleResponse(httpResp)
	return resp, err
}

// getByScopeCreateRequest creates the GetByScope request.
func (client *SystemAssignedIdentitiesClient) getByScopeCreateRequest(ctx context.Context, scope string, options *SystemAssignedIdentitiesClientGetByScopeOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.ManagedIdentity/identities/default"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-11-30")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getByScopeHandleResponse handles the GetByScope response.
func (client *SystemAssignedIdentitiesClient) getByScopeHandleResponse(resp *http.Response) (SystemAssignedIdentitiesClientGetByScopeResponse, error) {
	result := SystemAssignedIdentitiesClientGetByScopeResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SystemAssignedIdentity); err != nil {
		return SystemAssignedIdentitiesClientGetByScopeResponse{}, err
	}
	return result, nil
}
