//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

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
	"strconv"
	"strings"
)

// DraftPackagesClient contains the methods for the DraftPackages group.
// Don't use this type directly, use NewDraftPackagesClient() instead.
type DraftPackagesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewDraftPackagesClient creates a new instance of DraftPackagesClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewDraftPackagesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DraftPackagesClient, error) {
	cl, err := arm.NewClient(moduleName+".DraftPackagesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &DraftPackagesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCopyFromPackage - Copy package file and metadata from a package to this draft package
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - testBaseAccountName - The resource name of the Test Base Account.
//   - draftPackageName - The resource name of the Test Base Draft Package.
//   - parameters - Parameters supplied to the Test Base Draft Package CopyFromPackage operation.
//   - options - DraftPackagesClientBeginCopyFromPackageOptions contains the optional parameters for the DraftPackagesClient.BeginCopyFromPackage
//     method.
func (client *DraftPackagesClient) BeginCopyFromPackage(ctx context.Context, resourceGroupName string, testBaseAccountName string, draftPackageName string, parameters CopyFromPackageOperationParameters, options *DraftPackagesClientBeginCopyFromPackageOptions) (*runtime.Poller[DraftPackagesClientCopyFromPackageResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.copyFromPackage(ctx, resourceGroupName, testBaseAccountName, draftPackageName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[DraftPackagesClientCopyFromPackageResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[DraftPackagesClientCopyFromPackageResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CopyFromPackage - Copy package file and metadata from a package to this draft package
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01-preview
func (client *DraftPackagesClient) copyFromPackage(ctx context.Context, resourceGroupName string, testBaseAccountName string, draftPackageName string, parameters CopyFromPackageOperationParameters, options *DraftPackagesClientBeginCopyFromPackageOptions) (*http.Response, error) {
	req, err := client.copyFromPackageCreateRequest(ctx, resourceGroupName, testBaseAccountName, draftPackageName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// copyFromPackageCreateRequest creates the CopyFromPackage request.
func (client *DraftPackagesClient) copyFromPackageCreateRequest(ctx context.Context, resourceGroupName string, testBaseAccountName string, draftPackageName string, parameters CopyFromPackageOperationParameters, options *DraftPackagesClientBeginCopyFromPackageOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TestBase/testBaseAccounts/{testBaseAccountName}/draftPackages/{draftPackageName}/copyFromPackage"
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
	if draftPackageName == "" {
		return nil, errors.New("parameter draftPackageName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{draftPackageName}", url.PathEscape(draftPackageName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// Create - Creates or replaces a Test Base Draft Package.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - testBaseAccountName - The resource name of the Test Base Account.
//   - draftPackageName - The resource name of the Test Base Draft Package.
//   - parameters - Parameters supplied to create a Test Base Draft Package.
//   - options - DraftPackagesClientCreateOptions contains the optional parameters for the DraftPackagesClient.Create method.
func (client *DraftPackagesClient) Create(ctx context.Context, resourceGroupName string, testBaseAccountName string, draftPackageName string, parameters DraftPackageResource, options *DraftPackagesClientCreateOptions) (DraftPackagesClientCreateResponse, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, testBaseAccountName, draftPackageName, parameters, options)
	if err != nil {
		return DraftPackagesClientCreateResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DraftPackagesClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DraftPackagesClientCreateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createHandleResponse(resp)
}

// createCreateRequest creates the Create request.
func (client *DraftPackagesClient) createCreateRequest(ctx context.Context, resourceGroupName string, testBaseAccountName string, draftPackageName string, parameters DraftPackageResource, options *DraftPackagesClientCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TestBase/testBaseAccounts/{testBaseAccountName}/draftPackages/{draftPackageName}"
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
	if draftPackageName == "" {
		return nil, errors.New("parameter draftPackageName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{draftPackageName}", url.PathEscape(draftPackageName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createHandleResponse handles the Create response.
func (client *DraftPackagesClient) createHandleResponse(resp *http.Response) (DraftPackagesClientCreateResponse, error) {
	result := DraftPackagesClientCreateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DraftPackageResource); err != nil {
		return DraftPackagesClientCreateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes a Test Base Draft Package.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - testBaseAccountName - The resource name of the Test Base Account.
//   - draftPackageName - The resource name of the Test Base Draft Package.
//   - options - DraftPackagesClientDeleteOptions contains the optional parameters for the DraftPackagesClient.Delete method.
func (client *DraftPackagesClient) Delete(ctx context.Context, resourceGroupName string, testBaseAccountName string, draftPackageName string, options *DraftPackagesClientDeleteOptions) (DraftPackagesClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, testBaseAccountName, draftPackageName, options)
	if err != nil {
		return DraftPackagesClientDeleteResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DraftPackagesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return DraftPackagesClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return DraftPackagesClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *DraftPackagesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, testBaseAccountName string, draftPackageName string, options *DraftPackagesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TestBase/testBaseAccounts/{testBaseAccountName}/draftPackages/{draftPackageName}"
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
	if draftPackageName == "" {
		return nil, errors.New("parameter draftPackageName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{draftPackageName}", url.PathEscape(draftPackageName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginExtractFile - Performs extracting file operation for a Test Base Draft Package
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - testBaseAccountName - The resource name of the Test Base Account.
//   - draftPackageName - The resource name of the Test Base Draft Package.
//   - parameters - Parameters supplied to the Test Base Draft Package ExtractFile operation.
//   - options - DraftPackagesClientBeginExtractFileOptions contains the optional parameters for the DraftPackagesClient.BeginExtractFile
//     method.
func (client *DraftPackagesClient) BeginExtractFile(ctx context.Context, resourceGroupName string, testBaseAccountName string, draftPackageName string, parameters ExtractFileOperationParameters, options *DraftPackagesClientBeginExtractFileOptions) (*runtime.Poller[DraftPackagesClientExtractFileResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.extractFile(ctx, resourceGroupName, testBaseAccountName, draftPackageName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[DraftPackagesClientExtractFileResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[DraftPackagesClientExtractFileResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// ExtractFile - Performs extracting file operation for a Test Base Draft Package
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01-preview
func (client *DraftPackagesClient) extractFile(ctx context.Context, resourceGroupName string, testBaseAccountName string, draftPackageName string, parameters ExtractFileOperationParameters, options *DraftPackagesClientBeginExtractFileOptions) (*http.Response, error) {
	req, err := client.extractFileCreateRequest(ctx, resourceGroupName, testBaseAccountName, draftPackageName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// extractFileCreateRequest creates the ExtractFile request.
func (client *DraftPackagesClient) extractFileCreateRequest(ctx context.Context, resourceGroupName string, testBaseAccountName string, draftPackageName string, parameters ExtractFileOperationParameters, options *DraftPackagesClientBeginExtractFileOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TestBase/testBaseAccounts/{testBaseAccountName}/draftPackages/{draftPackageName}/extractFile"
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
	if draftPackageName == "" {
		return nil, errors.New("parameter draftPackageName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{draftPackageName}", url.PathEscape(draftPackageName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginGenerateFoldersAndScripts - Generates folders and scripts
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - testBaseAccountName - The resource name of the Test Base Account.
//   - draftPackageName - The resource name of the Test Base Draft Package.
//   - parameters - Parameters supplied to the Test Base Draft Package Generate operation.
//   - options - DraftPackagesClientBeginGenerateFoldersAndScriptsOptions contains the optional parameters for the DraftPackagesClient.BeginGenerateFoldersAndScripts
//     method.
func (client *DraftPackagesClient) BeginGenerateFoldersAndScripts(ctx context.Context, resourceGroupName string, testBaseAccountName string, draftPackageName string, parameters GenerateOperationParameters, options *DraftPackagesClientBeginGenerateFoldersAndScriptsOptions) (*runtime.Poller[DraftPackagesClientGenerateFoldersAndScriptsResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.generateFoldersAndScripts(ctx, resourceGroupName, testBaseAccountName, draftPackageName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[DraftPackagesClientGenerateFoldersAndScriptsResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[DraftPackagesClientGenerateFoldersAndScriptsResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// GenerateFoldersAndScripts - Generates folders and scripts
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01-preview
func (client *DraftPackagesClient) generateFoldersAndScripts(ctx context.Context, resourceGroupName string, testBaseAccountName string, draftPackageName string, parameters GenerateOperationParameters, options *DraftPackagesClientBeginGenerateFoldersAndScriptsOptions) (*http.Response, error) {
	req, err := client.generateFoldersAndScriptsCreateRequest(ctx, resourceGroupName, testBaseAccountName, draftPackageName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// generateFoldersAndScriptsCreateRequest creates the GenerateFoldersAndScripts request.
func (client *DraftPackagesClient) generateFoldersAndScriptsCreateRequest(ctx context.Context, resourceGroupName string, testBaseAccountName string, draftPackageName string, parameters GenerateOperationParameters, options *DraftPackagesClientBeginGenerateFoldersAndScriptsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TestBase/testBaseAccounts/{testBaseAccountName}/draftPackages/{draftPackageName}/generateFoldersAndScripts"
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
	if draftPackageName == "" {
		return nil, errors.New("parameter draftPackageName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{draftPackageName}", url.PathEscape(draftPackageName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// Get - Gets a Test Base Draft Package.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - testBaseAccountName - The resource name of the Test Base Account.
//   - draftPackageName - The resource name of the Test Base Draft Package.
//   - options - DraftPackagesClientGetOptions contains the optional parameters for the DraftPackagesClient.Get method.
func (client *DraftPackagesClient) Get(ctx context.Context, resourceGroupName string, testBaseAccountName string, draftPackageName string, options *DraftPackagesClientGetOptions) (DraftPackagesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, testBaseAccountName, draftPackageName, options)
	if err != nil {
		return DraftPackagesClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DraftPackagesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DraftPackagesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DraftPackagesClient) getCreateRequest(ctx context.Context, resourceGroupName string, testBaseAccountName string, draftPackageName string, options *DraftPackagesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TestBase/testBaseAccounts/{testBaseAccountName}/draftPackages/{draftPackageName}"
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
	if draftPackageName == "" {
		return nil, errors.New("parameter draftPackageName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{draftPackageName}", url.PathEscape(draftPackageName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DraftPackagesClient) getHandleResponse(resp *http.Response) (DraftPackagesClientGetResponse, error) {
	result := DraftPackagesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DraftPackageResource); err != nil {
		return DraftPackagesClientGetResponse{}, err
	}
	return result, nil
}

// GetPath - Gets draft package path and temp working path with SAS.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - testBaseAccountName - The resource name of the Test Base Account.
//   - draftPackageName - The resource name of the Test Base Draft Package.
//   - options - DraftPackagesClientGetPathOptions contains the optional parameters for the DraftPackagesClient.GetPath method.
func (client *DraftPackagesClient) GetPath(ctx context.Context, resourceGroupName string, testBaseAccountName string, draftPackageName string, options *DraftPackagesClientGetPathOptions) (DraftPackagesClientGetPathResponse, error) {
	req, err := client.getPathCreateRequest(ctx, resourceGroupName, testBaseAccountName, draftPackageName, options)
	if err != nil {
		return DraftPackagesClientGetPathResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DraftPackagesClientGetPathResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DraftPackagesClientGetPathResponse{}, runtime.NewResponseError(resp)
	}
	return client.getPathHandleResponse(resp)
}

// getPathCreateRequest creates the GetPath request.
func (client *DraftPackagesClient) getPathCreateRequest(ctx context.Context, resourceGroupName string, testBaseAccountName string, draftPackageName string, options *DraftPackagesClientGetPathOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TestBase/testBaseAccounts/{testBaseAccountName}/draftPackages/{draftPackageName}/getPath"
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
	if draftPackageName == "" {
		return nil, errors.New("parameter draftPackageName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{draftPackageName}", url.PathEscape(draftPackageName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getPathHandleResponse handles the GetPath response.
func (client *DraftPackagesClient) getPathHandleResponse(resp *http.Response) (DraftPackagesClientGetPathResponse, error) {
	result := DraftPackagesClientGetPathResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DraftPackageGetPathResponse); err != nil {
		return DraftPackagesClientGetPathResponse{}, err
	}
	return result, nil
}

// NewListByTestBaseAccountPager - Lists all the draft packages under a Test Base Account.
//
// Generated from API version 2023-07-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - testBaseAccountName - The resource name of the Test Base Account.
//   - options - DraftPackagesClientListByTestBaseAccountOptions contains the optional parameters for the DraftPackagesClient.NewListByTestBaseAccountPager
//     method.
func (client *DraftPackagesClient) NewListByTestBaseAccountPager(resourceGroupName string, testBaseAccountName string, options *DraftPackagesClientListByTestBaseAccountOptions) *runtime.Pager[DraftPackagesClientListByTestBaseAccountResponse] {
	return runtime.NewPager(runtime.PagingHandler[DraftPackagesClientListByTestBaseAccountResponse]{
		More: func(page DraftPackagesClientListByTestBaseAccountResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DraftPackagesClientListByTestBaseAccountResponse) (DraftPackagesClientListByTestBaseAccountResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByTestBaseAccountCreateRequest(ctx, resourceGroupName, testBaseAccountName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DraftPackagesClientListByTestBaseAccountResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return DraftPackagesClientListByTestBaseAccountResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DraftPackagesClientListByTestBaseAccountResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByTestBaseAccountHandleResponse(resp)
		},
	})
}

// listByTestBaseAccountCreateRequest creates the ListByTestBaseAccount request.
func (client *DraftPackagesClient) listByTestBaseAccountCreateRequest(ctx context.Context, resourceGroupName string, testBaseAccountName string, options *DraftPackagesClientListByTestBaseAccountOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TestBase/testBaseAccounts/{testBaseAccountName}/draftPackages"
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.PackageName != nil {
		reqQP.Set("packageName", *options.PackageName)
	}
	if options != nil && options.EditPackage != nil {
		reqQP.Set("editPackage", strconv.FormatBool(*options.EditPackage))
	}
	reqQP.Set("api-version", "2023-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByTestBaseAccountHandleResponse handles the ListByTestBaseAccount response.
func (client *DraftPackagesClient) listByTestBaseAccountHandleResponse(resp *http.Response) (DraftPackagesClientListByTestBaseAccountResponse, error) {
	result := DraftPackagesClientListByTestBaseAccountResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DraftPackageListResult); err != nil {
		return DraftPackagesClientListByTestBaseAccountResponse{}, err
	}
	return result, nil
}

// Update - Updates an existing Test Base Draft Package.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - testBaseAccountName - The resource name of the Test Base Account.
//   - draftPackageName - The resource name of the Test Base Draft Package.
//   - parameters - Parameters supplied to update a Test Base Draft Package.
//   - options - DraftPackagesClientUpdateOptions contains the optional parameters for the DraftPackagesClient.Update method.
func (client *DraftPackagesClient) Update(ctx context.Context, resourceGroupName string, testBaseAccountName string, draftPackageName string, parameters DraftPackageUpdateParameters, options *DraftPackagesClientUpdateOptions) (DraftPackagesClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, testBaseAccountName, draftPackageName, parameters, options)
	if err != nil {
		return DraftPackagesClientUpdateResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DraftPackagesClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DraftPackagesClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *DraftPackagesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, testBaseAccountName string, draftPackageName string, parameters DraftPackageUpdateParameters, options *DraftPackagesClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TestBase/testBaseAccounts/{testBaseAccountName}/draftPackages/{draftPackageName}"
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
	if draftPackageName == "" {
		return nil, errors.New("parameter draftPackageName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{draftPackageName}", url.PathEscape(draftPackageName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateHandleResponse handles the Update response.
func (client *DraftPackagesClient) updateHandleResponse(resp *http.Response) (DraftPackagesClientUpdateResponse, error) {
	result := DraftPackagesClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DraftPackageResource); err != nil {
		return DraftPackagesClientUpdateResponse{}, err
	}
	return result, nil
}
