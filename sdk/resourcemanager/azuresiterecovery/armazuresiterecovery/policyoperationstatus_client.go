//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armazuresiterecovery

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

// PolicyOperationStatusClient contains the methods for the PolicyOperationStatus group.
// Don't use this type directly, use NewPolicyOperationStatusClient() instead.
type PolicyOperationStatusClient struct {
	internal *arm.Client
}

// NewPolicyOperationStatusClient creates a new instance of PolicyOperationStatusClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewPolicyOperationStatusClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*PolicyOperationStatusClient, error) {
	cl, err := arm.NewClient(moduleName+".PolicyOperationStatusClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &PolicyOperationStatusClient{
		internal: cl,
	}
	return client, nil
}

// Get - Tracks the results of an asynchronous operation on the policy.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-02-16-preview
//   - subscriptionID - The subscription Id.
//   - resourceGroupName - Resource group name.
//   - vaultName - Vault name.
//   - policyName - Policy name.
//   - operationID - Operation Id.
//   - options - PolicyOperationStatusClientGetOptions contains the optional parameters for the PolicyOperationStatusClient.Get
//     method.
func (client *PolicyOperationStatusClient) Get(ctx context.Context, subscriptionID string, resourceGroupName string, vaultName string, policyName string, operationID string, options *PolicyOperationStatusClientGetOptions) (PolicyOperationStatusClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, subscriptionID, resourceGroupName, vaultName, policyName, operationID, options)
	if err != nil {
		return PolicyOperationStatusClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PolicyOperationStatusClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PolicyOperationStatusClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *PolicyOperationStatusClient) getCreateRequest(ctx context.Context, subscriptionID string, resourceGroupName string, vaultName string, policyName string, operationID string, options *PolicyOperationStatusClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataReplication/replicationVaults/{vaultName}/replicationPolicies/{policyName}/operations/{operationId}"
	if subscriptionID == "" {
		return nil, errors.New("parameter subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vaultName == "" {
		return nil, errors.New("parameter vaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	if policyName == "" {
		return nil, errors.New("parameter policyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{policyName}", url.PathEscape(policyName))
	if operationID == "" {
		return nil, errors.New("parameter operationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{operationId}", url.PathEscape(operationID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-16-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PolicyOperationStatusClient) getHandleResponse(resp *http.Response) (PolicyOperationStatusClientGetResponse, error) {
	result := PolicyOperationStatusClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OperationStatus); err != nil {
		return PolicyOperationStatusClientGetResponse{}, err
	}
	return result, nil
}
