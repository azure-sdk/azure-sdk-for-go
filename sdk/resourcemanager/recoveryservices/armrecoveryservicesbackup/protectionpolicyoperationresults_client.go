// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armrecoveryservicesbackup

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

// ProtectionPolicyOperationResultsClient contains the methods for the ProtectionPolicyOperationResults group.
// Don't use this type directly, use NewProtectionPolicyOperationResultsClient() instead.
type ProtectionPolicyOperationResultsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewProtectionPolicyOperationResultsClient creates a new instance of ProtectionPolicyOperationResultsClient with the specified values.
//   - subscriptionID - The subscription Id.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewProtectionPolicyOperationResultsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ProtectionPolicyOperationResultsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ProtectionPolicyOperationResultsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Provides the result of an operation.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-11-01-preview
//   - vaultName - The name of the recovery services vault.
//   - resourceGroupName - The name of the resource group where the recovery services vault is present.
//   - policyName - Backup policy name whose operation's result needs to be fetched.
//   - operationID - Operation ID which represents the operation whose result needs to be fetched.
//   - options - ProtectionPolicyOperationResultsClientGetOptions contains the optional parameters for the ProtectionPolicyOperationResultsClient.Get
//     method.
func (client *ProtectionPolicyOperationResultsClient) Get(ctx context.Context, vaultName string, resourceGroupName string, policyName string, operationID string, options *ProtectionPolicyOperationResultsClientGetOptions) (ProtectionPolicyOperationResultsClientGetResponse, error) {
	var err error
	const operationName = "ProtectionPolicyOperationResultsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, vaultName, resourceGroupName, policyName, operationID, options)
	if err != nil {
		return ProtectionPolicyOperationResultsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ProtectionPolicyOperationResultsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ProtectionPolicyOperationResultsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ProtectionPolicyOperationResultsClient) getCreateRequest(ctx context.Context, vaultName string, resourceGroupName string, policyName string, operationID string, _ *ProtectionPolicyOperationResultsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupPolicies/{policyName}/operationResults/{operationId}"
	if vaultName == "" {
		return nil, errors.New("parameter vaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
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
	reqQP.Set("api-version", "2024-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ProtectionPolicyOperationResultsClient) getHandleResponse(resp *http.Response) (ProtectionPolicyOperationResultsClientGetResponse, error) {
	result := ProtectionPolicyOperationResultsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProtectionPolicyResource); err != nil {
		return ProtectionPolicyOperationResultsClientGetResponse{}, err
	}
	return result, nil
}
