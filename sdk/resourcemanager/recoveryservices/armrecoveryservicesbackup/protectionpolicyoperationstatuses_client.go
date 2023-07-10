//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

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

// ProtectionPolicyOperationStatusesClient contains the methods for the ProtectionPolicyOperationStatuses group.
// Don't use this type directly, use NewProtectionPolicyOperationStatusesClient() instead.
type ProtectionPolicyOperationStatusesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewProtectionPolicyOperationStatusesClient creates a new instance of ProtectionPolicyOperationStatusesClient with the specified values.
//   - subscriptionID - The subscription Id.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewProtectionPolicyOperationStatusesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ProtectionPolicyOperationStatusesClient, error) {
	cl, err := arm.NewClient(moduleName+".ProtectionPolicyOperationStatusesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ProtectionPolicyOperationStatusesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Provides the status of the asynchronous operations like backup, restore. The status can be in progress, completed
// or failed. You can refer to the Operation Status enum for all the possible states of
// an operation. Some operations create jobs. This method returns the list of jobs associated with operation.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01
//   - vaultName - The name of the recovery services vault.
//   - resourceGroupName - The name of the resource group where the recovery services vault is present.
//   - policyName - Backup policy name whose operation's status needs to be fetched.
//   - operationID - Operation ID which represents an operation whose status needs to be fetched.
//   - options - ProtectionPolicyOperationStatusesClientGetOptions contains the optional parameters for the ProtectionPolicyOperationStatusesClient.Get
//     method.
func (client *ProtectionPolicyOperationStatusesClient) Get(ctx context.Context, vaultName string, resourceGroupName string, policyName string, operationID string, options *ProtectionPolicyOperationStatusesClientGetOptions) (ProtectionPolicyOperationStatusesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, vaultName, resourceGroupName, policyName, operationID, options)
	if err != nil {
		return ProtectionPolicyOperationStatusesClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ProtectionPolicyOperationStatusesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ProtectionPolicyOperationStatusesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ProtectionPolicyOperationStatusesClient) getCreateRequest(ctx context.Context, vaultName string, resourceGroupName string, policyName string, operationID string, options *ProtectionPolicyOperationStatusesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupPolicies/{policyName}/operations/{operationId}"
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
	reqQP.Set("api-version", "2023-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ProtectionPolicyOperationStatusesClient) getHandleResponse(resp *http.Response) (ProtectionPolicyOperationStatusesClientGetResponse, error) {
	result := ProtectionPolicyOperationStatusesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OperationStatus); err != nil {
		return ProtectionPolicyOperationStatusesClientGetResponse{}, err
	}
	return result, nil
}
