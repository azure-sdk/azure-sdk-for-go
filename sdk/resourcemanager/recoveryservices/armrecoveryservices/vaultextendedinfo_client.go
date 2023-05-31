//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armrecoveryservices

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

// VaultExtendedInfoClient contains the methods for the VaultExtendedInfo group.
// Don't use this type directly, use NewVaultExtendedInfoClient() instead.
type VaultExtendedInfoClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewVaultExtendedInfoClient creates a new instance of VaultExtendedInfoClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewVaultExtendedInfoClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*VaultExtendedInfoClient, error) {
	cl, err := arm.NewClient(moduleName+".VaultExtendedInfoClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &VaultExtendedInfoClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Create vault extended info.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - vaultName - The name of the recovery services vault.
//   - resourceExtendedInfoDetails - Details of ResourceExtendedInfo
//   - options - VaultExtendedInfoClientCreateOrUpdateOptions contains the optional parameters for the VaultExtendedInfoClient.CreateOrUpdate
//     method.
func (client *VaultExtendedInfoClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, vaultName string, resourceExtendedInfoDetails VaultExtendedInfoResource, options *VaultExtendedInfoClientCreateOrUpdateOptions) (VaultExtendedInfoClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, vaultName, resourceExtendedInfoDetails, options)
	if err != nil {
		return VaultExtendedInfoClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return VaultExtendedInfoClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VaultExtendedInfoClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *VaultExtendedInfoClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, vaultName string, resourceExtendedInfoDetails VaultExtendedInfoResource, options *VaultExtendedInfoClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/extendedInformation/vaultExtendedInfo"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vaultName == "" {
		return nil, errors.New("parameter vaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, resourceExtendedInfoDetails)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *VaultExtendedInfoClient) createOrUpdateHandleResponse(resp *http.Response) (VaultExtendedInfoClientCreateOrUpdateResponse, error) {
	result := VaultExtendedInfoClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VaultExtendedInfoResource); err != nil {
		return VaultExtendedInfoClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Get - Get the vault extended info.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - vaultName - The name of the recovery services vault.
//   - options - VaultExtendedInfoClientGetOptions contains the optional parameters for the VaultExtendedInfoClient.Get method.
func (client *VaultExtendedInfoClient) Get(ctx context.Context, resourceGroupName string, vaultName string, options *VaultExtendedInfoClientGetOptions) (VaultExtendedInfoClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, vaultName, options)
	if err != nil {
		return VaultExtendedInfoClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return VaultExtendedInfoClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VaultExtendedInfoClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *VaultExtendedInfoClient) getCreateRequest(ctx context.Context, resourceGroupName string, vaultName string, options *VaultExtendedInfoClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/extendedInformation/vaultExtendedInfo"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vaultName == "" {
		return nil, errors.New("parameter vaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
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
func (client *VaultExtendedInfoClient) getHandleResponse(resp *http.Response) (VaultExtendedInfoClientGetResponse, error) {
	result := VaultExtendedInfoClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VaultExtendedInfoResource); err != nil {
		return VaultExtendedInfoClientGetResponse{}, err
	}
	return result, nil
}

// Update - Update vault extended info.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - vaultName - The name of the recovery services vault.
//   - resourceExtendedInfoDetails - Details of ResourceExtendedInfo
//   - options - VaultExtendedInfoClientUpdateOptions contains the optional parameters for the VaultExtendedInfoClient.Update
//     method.
func (client *VaultExtendedInfoClient) Update(ctx context.Context, resourceGroupName string, vaultName string, resourceExtendedInfoDetails VaultExtendedInfoResource, options *VaultExtendedInfoClientUpdateOptions) (VaultExtendedInfoClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, vaultName, resourceExtendedInfoDetails, options)
	if err != nil {
		return VaultExtendedInfoClientUpdateResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return VaultExtendedInfoClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VaultExtendedInfoClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *VaultExtendedInfoClient) updateCreateRequest(ctx context.Context, resourceGroupName string, vaultName string, resourceExtendedInfoDetails VaultExtendedInfoResource, options *VaultExtendedInfoClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/extendedInformation/vaultExtendedInfo"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vaultName == "" {
		return nil, errors.New("parameter vaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, resourceExtendedInfoDetails)
}

// updateHandleResponse handles the Update response.
func (client *VaultExtendedInfoClient) updateHandleResponse(resp *http.Response) (VaultExtendedInfoClientUpdateResponse, error) {
	result := VaultExtendedInfoClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VaultExtendedInfoResource); err != nil {
		return VaultExtendedInfoClientUpdateResponse{}, err
	}
	return result, nil
}
