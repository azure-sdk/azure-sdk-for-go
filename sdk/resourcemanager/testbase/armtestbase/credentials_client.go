// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armtestbase

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

// CredentialsClient contains the methods for the Credentials group.
// Don't use this type directly, use NewCredentialsClient() instead.
type CredentialsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewCredentialsClient creates a new instance of CredentialsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewCredentialsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*CredentialsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &CredentialsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Create - Creates or replaces a Test Base Credential.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - testBaseAccountName - The resource name of the Test Base Account.
//   - credentialName - The credential resource name.
//   - parameters - Parameters supplied to create a Test Base Credential.
//   - options - CredentialsClientCreateOptions contains the optional parameters for the CredentialsClient.Create method.
func (client *CredentialsClient) Create(ctx context.Context, resourceGroupName string, testBaseAccountName string, credentialName string, parameters CredentialResource, options *CredentialsClientCreateOptions) (CredentialsClientCreateResponse, error) {
	var err error
	const operationName = "CredentialsClient.Create"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createCreateRequest(ctx, resourceGroupName, testBaseAccountName, credentialName, parameters, options)
	if err != nil {
		return CredentialsClientCreateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return CredentialsClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return CredentialsClientCreateResponse{}, err
	}
	resp, err := client.createHandleResponse(httpResp)
	return resp, err
}

// createCreateRequest creates the Create request.
func (client *CredentialsClient) createCreateRequest(ctx context.Context, resourceGroupName string, testBaseAccountName string, credentialName string, parameters CredentialResource, _ *CredentialsClientCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TestBase/testBaseAccounts/{testBaseAccountName}/credentials/{credentialName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if testBaseAccountName == "" {
		return nil, errors.New("parameter testBaseAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{testBaseAccountName}", url.PathEscape(testBaseAccountName))
	if credentialName == "" {
		return nil, errors.New("parameter credentialName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{credentialName}", url.PathEscape(credentialName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// createHandleResponse handles the Create response.
func (client *CredentialsClient) createHandleResponse(resp *http.Response) (CredentialsClientCreateResponse, error) {
	result := CredentialsClientCreateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CredentialResource); err != nil {
		return CredentialsClientCreateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes an existing test base credential.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - testBaseAccountName - The resource name of the Test Base Account.
//   - credentialName - The credential resource name.
//   - options - CredentialsClientDeleteOptions contains the optional parameters for the CredentialsClient.Delete method.
func (client *CredentialsClient) Delete(ctx context.Context, resourceGroupName string, testBaseAccountName string, credentialName string, options *CredentialsClientDeleteOptions) (CredentialsClientDeleteResponse, error) {
	var err error
	const operationName = "CredentialsClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, testBaseAccountName, credentialName, options)
	if err != nil {
		return CredentialsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return CredentialsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return CredentialsClientDeleteResponse{}, err
	}
	return CredentialsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *CredentialsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, testBaseAccountName string, credentialName string, _ *CredentialsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TestBase/testBaseAccounts/{testBaseAccountName}/credentials/{credentialName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if testBaseAccountName == "" {
		return nil, errors.New("parameter testBaseAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{testBaseAccountName}", url.PathEscape(testBaseAccountName))
	if credentialName == "" {
		return nil, errors.New("parameter credentialName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{credentialName}", url.PathEscape(credentialName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Update - Updates an existing test base credential.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - testBaseAccountName - The resource name of the Test Base Account.
//   - credentialName - The credential resource name.
//   - parameters - Parameters supplied to create a test base credential.
//   - options - CredentialsClientUpdateOptions contains the optional parameters for the CredentialsClient.Update method.
func (client *CredentialsClient) Update(ctx context.Context, resourceGroupName string, testBaseAccountName string, credentialName string, parameters CredentialResource, options *CredentialsClientUpdateOptions) (CredentialsClientUpdateResponse, error) {
	var err error
	const operationName = "CredentialsClient.Update"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, testBaseAccountName, credentialName, parameters, options)
	if err != nil {
		return CredentialsClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return CredentialsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return CredentialsClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *CredentialsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, testBaseAccountName string, credentialName string, parameters CredentialResource, _ *CredentialsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TestBase/testBaseAccounts/{testBaseAccountName}/credentials/{credentialName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if testBaseAccountName == "" {
		return nil, errors.New("parameter testBaseAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{testBaseAccountName}", url.PathEscape(testBaseAccountName))
	if credentialName == "" {
		return nil, errors.New("parameter credentialName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{credentialName}", url.PathEscape(credentialName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *CredentialsClient) updateHandleResponse(resp *http.Response) (CredentialsClientUpdateResponse, error) {
	result := CredentialsClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CredentialResource); err != nil {
		return CredentialsClientUpdateResponse{}, err
	}
	return result, nil
}
