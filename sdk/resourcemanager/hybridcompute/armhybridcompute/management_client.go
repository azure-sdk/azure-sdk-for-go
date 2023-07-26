//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

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

// ManagementClient contains the methods for the HybridComputeManagementClient group.
// Don't use this type directly, use NewManagementClient() instead.
type ManagementClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewManagementClient creates a new instance of ManagementClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewManagementClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ManagementClient, error) {
	cl, err := arm.NewClient(moduleName+".ManagementClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ManagementClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginUpgradeExtensions - The operation to Upgrade Machine Extensions.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-15-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - machineName - The name of the hybrid machine.
//   - extensionUpgradeParameters - Parameters supplied to the Upgrade Extensions operation.
//   - options - ManagementClientBeginUpgradeExtensionsOptions contains the optional parameters for the ManagementClient.BeginUpgradeExtensions
//     method.
func (client *ManagementClient) BeginUpgradeExtensions(ctx context.Context, resourceGroupName string, machineName string, extensionUpgradeParameters MachineExtensionUpgrade, options *ManagementClientBeginUpgradeExtensionsOptions) (*runtime.Poller[ManagementClientUpgradeExtensionsResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.upgradeExtensions(ctx, resourceGroupName, machineName, extensionUpgradeParameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[ManagementClientUpgradeExtensionsResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[ManagementClientUpgradeExtensionsResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// UpgradeExtensions - The operation to Upgrade Machine Extensions.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-15-preview
func (client *ManagementClient) upgradeExtensions(ctx context.Context, resourceGroupName string, machineName string, extensionUpgradeParameters MachineExtensionUpgrade, options *ManagementClientBeginUpgradeExtensionsOptions) (*http.Response, error) {
	var err error
	req, err := client.upgradeExtensionsCreateRequest(ctx, resourceGroupName, machineName, extensionUpgradeParameters, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// upgradeExtensionsCreateRequest creates the UpgradeExtensions request.
func (client *ManagementClient) upgradeExtensionsCreateRequest(ctx context.Context, resourceGroupName string, machineName string, extensionUpgradeParameters MachineExtensionUpgrade, options *ManagementClientBeginUpgradeExtensionsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridCompute/machines/{machineName}/upgradeExtensions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if machineName == "" {
		return nil, errors.New("parameter machineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{machineName}", url.PathEscape(machineName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, extensionUpgradeParameters); err != nil {
		return nil, err
	}
	return req, nil
}
