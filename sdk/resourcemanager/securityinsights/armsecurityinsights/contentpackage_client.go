//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsecurityinsights

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

// ContentPackageClient contains the methods for the ContentPackage group.
// Don't use this type directly, use NewContentPackageClient() instead.
type ContentPackageClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewContentPackageClient creates a new instance of ContentPackageClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewContentPackageClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ContentPackageClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ContentPackageClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Install - Install a package to the workspace.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - packageID - package Id
//   - packageInstallationProperties - Package installation properties
//   - options - ContentPackageClientInstallOptions contains the optional parameters for the ContentPackageClient.Install method.
func (client *ContentPackageClient) Install(ctx context.Context, resourceGroupName string, workspaceName string, packageID string, packageInstallationProperties PackageModel, options *ContentPackageClientInstallOptions) (ContentPackageClientInstallResponse, error) {
	var err error
	const operationName = "ContentPackageClient.Install"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.installCreateRequest(ctx, resourceGroupName, workspaceName, packageID, packageInstallationProperties, options)
	if err != nil {
		return ContentPackageClientInstallResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ContentPackageClientInstallResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return ContentPackageClientInstallResponse{}, err
	}
	resp, err := client.installHandleResponse(httpResp)
	return resp, err
}

// installCreateRequest creates the Install request.
func (client *ContentPackageClient) installCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, packageID string, packageInstallationProperties PackageModel, options *ContentPackageClientInstallOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/contentPackages/{packageId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if packageID == "" {
		return nil, errors.New("parameter packageID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{packageId}", url.PathEscape(packageID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, packageInstallationProperties); err != nil {
		return nil, err
	}
	return req, nil
}

// installHandleResponse handles the Install response.
func (client *ContentPackageClient) installHandleResponse(resp *http.Response) (ContentPackageClientInstallResponse, error) {
	result := ContentPackageClientInstallResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PackageModel); err != nil {
		return ContentPackageClientInstallResponse{}, err
	}
	return result, nil
}

// Uninstall - Uninstall a package from the workspace.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - packageID - package Id
//   - options - ContentPackageClientUninstallOptions contains the optional parameters for the ContentPackageClient.Uninstall
//     method.
func (client *ContentPackageClient) Uninstall(ctx context.Context, resourceGroupName string, workspaceName string, packageID string, options *ContentPackageClientUninstallOptions) (ContentPackageClientUninstallResponse, error) {
	var err error
	const operationName = "ContentPackageClient.Uninstall"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.uninstallCreateRequest(ctx, resourceGroupName, workspaceName, packageID, options)
	if err != nil {
		return ContentPackageClientUninstallResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ContentPackageClientUninstallResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ContentPackageClientUninstallResponse{}, err
	}
	return ContentPackageClientUninstallResponse{}, nil
}

// uninstallCreateRequest creates the Uninstall request.
func (client *ContentPackageClient) uninstallCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, packageID string, options *ContentPackageClientUninstallOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/contentPackages/{packageId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if packageID == "" {
		return nil, errors.New("parameter packageID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{packageId}", url.PathEscape(packageID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}
