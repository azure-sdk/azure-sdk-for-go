//go:build go1.18
// +build go1.18

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

// ProtectionIntentClient contains the methods for the ProtectionIntent group.
// Don't use this type directly, use NewProtectionIntentClient() instead.
type ProtectionIntentClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewProtectionIntentClient creates a new instance of ProtectionIntentClient with the specified values.
//   - subscriptionID - The subscription Id.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewProtectionIntentClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ProtectionIntentClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ProtectionIntentClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Create Intent for Enabling backup of an item. This is a synchronous operation.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
//   - vaultName - The name of the recovery services vault.
//   - resourceGroupName - The name of the resource group where the recovery services vault is present.
//   - fabricName - Fabric name associated with the backup item.
//   - intentObjectName - Intent object name.
//   - parameters - resource backed up item
//   - options - ProtectionIntentClientCreateOrUpdateOptions contains the optional parameters for the ProtectionIntentClient.CreateOrUpdate
//     method.
func (client *ProtectionIntentClient) CreateOrUpdate(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, intentObjectName string, parameters ProtectionIntentResource, options *ProtectionIntentClientCreateOrUpdateOptions) (ProtectionIntentClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "ProtectionIntentClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, vaultName, resourceGroupName, fabricName, intentObjectName, parameters, options)
	if err != nil {
		return ProtectionIntentClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ProtectionIntentClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ProtectionIntentClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ProtectionIntentClient) createOrUpdateCreateRequest(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, intentObjectName string, parameters ProtectionIntentResource, options *ProtectionIntentClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupFabrics/{fabricName}/backupProtectionIntent/{intentObjectName}"
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
	if fabricName == "" {
		return nil, errors.New("parameter fabricName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fabricName}", url.PathEscape(fabricName))
	if intentObjectName == "" {
		return nil, errors.New("parameter intentObjectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{intentObjectName}", url.PathEscape(intentObjectName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ProtectionIntentClient) createOrUpdateHandleResponse(resp *http.Response) (ProtectionIntentClientCreateOrUpdateResponse, error) {
	result := ProtectionIntentClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProtectionIntentResource); err != nil {
		return ProtectionIntentClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Used to remove intent from an item
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
//   - vaultName - The name of the recovery services vault.
//   - resourceGroupName - The name of the resource group where the recovery services vault is present.
//   - fabricName - Fabric name associated with the intent.
//   - intentObjectName - Intent to be deleted.
//   - options - ProtectionIntentClientDeleteOptions contains the optional parameters for the ProtectionIntentClient.Delete method.
func (client *ProtectionIntentClient) Delete(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, intentObjectName string, options *ProtectionIntentClientDeleteOptions) (ProtectionIntentClientDeleteResponse, error) {
	var err error
	const operationName = "ProtectionIntentClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, vaultName, resourceGroupName, fabricName, intentObjectName, options)
	if err != nil {
		return ProtectionIntentClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ProtectionIntentClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ProtectionIntentClientDeleteResponse{}, err
	}
	return ProtectionIntentClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ProtectionIntentClient) deleteCreateRequest(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, intentObjectName string, options *ProtectionIntentClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupFabrics/{fabricName}/backupProtectionIntent/{intentObjectName}"
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
	if fabricName == "" {
		return nil, errors.New("parameter fabricName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fabricName}", url.PathEscape(fabricName))
	if intentObjectName == "" {
		return nil, errors.New("parameter intentObjectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{intentObjectName}", url.PathEscape(intentObjectName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Provides the details of the protection intent up item. This is an asynchronous operation. To know the status of the
// operation, call the GetItemOperationResult API.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
//   - vaultName - The name of the recovery services vault.
//   - resourceGroupName - The name of the resource group where the recovery services vault is present.
//   - fabricName - Fabric name associated with the backed up item.
//   - intentObjectName - Backed up item name whose details are to be fetched.
//   - options - ProtectionIntentClientGetOptions contains the optional parameters for the ProtectionIntentClient.Get method.
func (client *ProtectionIntentClient) Get(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, intentObjectName string, options *ProtectionIntentClientGetOptions) (ProtectionIntentClientGetResponse, error) {
	var err error
	const operationName = "ProtectionIntentClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, vaultName, resourceGroupName, fabricName, intentObjectName, options)
	if err != nil {
		return ProtectionIntentClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ProtectionIntentClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ProtectionIntentClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ProtectionIntentClient) getCreateRequest(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, intentObjectName string, options *ProtectionIntentClientGetOptions) (*policy.Request, error) {
	urlPath := "/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupFabrics/{fabricName}/backupProtectionIntent/{intentObjectName}"
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
	if fabricName == "" {
		return nil, errors.New("parameter fabricName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fabricName}", url.PathEscape(fabricName))
	if intentObjectName == "" {
		return nil, errors.New("parameter intentObjectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{intentObjectName}", url.PathEscape(intentObjectName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ProtectionIntentClient) getHandleResponse(resp *http.Response) (ProtectionIntentClientGetResponse, error) {
	result := ProtectionIntentClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProtectionIntentResource); err != nil {
		return ProtectionIntentClientGetResponse{}, err
	}
	return result, nil
}

// Validate - It will validate followings
// 1. Vault capacity
// 2. VM is already protected
// 3. Any VM related configuration passed in properties.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
//   - azureRegion - Azure region to hit Api
//   - parameters - Enable backup validation request on Virtual Machine
//   - options - ProtectionIntentClientValidateOptions contains the optional parameters for the ProtectionIntentClient.Validate
//     method.
func (client *ProtectionIntentClient) Validate(ctx context.Context, azureRegion string, parameters PreValidateEnableBackupRequest, options *ProtectionIntentClientValidateOptions) (ProtectionIntentClientValidateResponse, error) {
	var err error
	const operationName = "ProtectionIntentClient.Validate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.validateCreateRequest(ctx, azureRegion, parameters, options)
	if err != nil {
		return ProtectionIntentClientValidateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ProtectionIntentClientValidateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ProtectionIntentClientValidateResponse{}, err
	}
	resp, err := client.validateHandleResponse(httpResp)
	return resp, err
}

// validateCreateRequest creates the Validate request.
func (client *ProtectionIntentClient) validateCreateRequest(ctx context.Context, azureRegion string, parameters PreValidateEnableBackupRequest, options *ProtectionIntentClientValidateOptions) (*policy.Request, error) {
	urlPath := "/Subscriptions/{subscriptionId}/providers/Microsoft.RecoveryServices/locations/{azureRegion}/backupPreValidateProtection"
	if azureRegion == "" {
		return nil, errors.New("parameter azureRegion cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{azureRegion}", url.PathEscape(azureRegion))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// validateHandleResponse handles the Validate response.
func (client *ProtectionIntentClient) validateHandleResponse(resp *http.Response) (ProtectionIntentClientValidateResponse, error) {
	result := ProtectionIntentClientValidateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PreValidateEnableBackupResponse); err != nil {
		return ProtectionIntentClientValidateResponse{}, err
	}
	return result, nil
}
