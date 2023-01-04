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
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// RegisteredIdentitiesClient contains the methods for the RegisteredIdentities group.
// Don't use this type directly, use NewRegisteredIdentitiesClient() instead.
type RegisteredIdentitiesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewRegisteredIdentitiesClient creates a new instance of RegisteredIdentitiesClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewRegisteredIdentitiesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*RegisteredIdentitiesClient, error) {
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
	client := &RegisteredIdentitiesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Delete - Unregisters the given container from your Recovery Services vault.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2023-01-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// vaultName - The name of the recovery services vault.
// identityName - Name of the protection container to unregister.
// options - RegisteredIdentitiesClientDeleteOptions contains the optional parameters for the RegisteredIdentitiesClient.Delete
// method.
func (client *RegisteredIdentitiesClient) Delete(ctx context.Context, resourceGroupName string, vaultName string, identityName string, options *RegisteredIdentitiesClientDeleteOptions) (RegisteredIdentitiesClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, vaultName, identityName, options)
	if err != nil {
		return RegisteredIdentitiesClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RegisteredIdentitiesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusNoContent) {
		return RegisteredIdentitiesClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return RegisteredIdentitiesClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *RegisteredIdentitiesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, vaultName string, identityName string, options *RegisteredIdentitiesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/registeredIdentities/{identityName}"
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
	if identityName == "" {
		return nil, errors.New("parameter identityName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{identityName}", url.PathEscape(identityName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}
