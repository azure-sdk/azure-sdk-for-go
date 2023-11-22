//go:build go1.18
// +build go1.18

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

// PackagesClient contains the methods for the Packages group.
// Don't use this type directly, use NewPackagesClient() instead.
type PackagesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewPackagesClient creates a new instance of PackagesClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewPackagesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*PackagesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &PackagesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreate - Create or replace (overwrite/recreate, with potential downtime) a Test Base Package.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - testBaseAccountName - The resource name of the Test Base Account.
//   - packageName - The resource name of the Test Base Package.
//   - parameters - Parameters supplied to create a Test Base Package.
//   - options - PackagesClientBeginCreateOptions contains the optional parameters for the PackagesClient.BeginCreate method.
func (client *PackagesClient) BeginCreate(ctx context.Context, resourceGroupName string, testBaseAccountName string, packageName string, parameters PackageResource, options *PackagesClientBeginCreateOptions) (*runtime.Poller[PackagesClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, testBaseAccountName, packageName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[PackagesClientCreateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[PackagesClientCreateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Create - Create or replace (overwrite/recreate, with potential downtime) a Test Base Package.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-01-preview
func (client *PackagesClient) create(ctx context.Context, resourceGroupName string, testBaseAccountName string, packageName string, parameters PackageResource, options *PackagesClientBeginCreateOptions) (*http.Response, error) {
	var err error
	const operationName = "PackagesClient.BeginCreate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createCreateRequest(ctx, resourceGroupName, testBaseAccountName, packageName, parameters, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createCreateRequest creates the Create request.
func (client *PackagesClient) createCreateRequest(ctx context.Context, resourceGroupName string, testBaseAccountName string, packageName string, parameters PackageResource, options *PackagesClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TestBase/testBaseAccounts/{testBaseAccountName}/packages/{packageName}"
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
	if packageName == "" {
		return nil, errors.New("parameter packageName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{packageName}", url.PathEscape(packageName))
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

// BeginDelete - Deletes a Test Base Package.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - testBaseAccountName - The resource name of the Test Base Account.
//   - packageName - The resource name of the Test Base Package.
//   - options - PackagesClientBeginDeleteOptions contains the optional parameters for the PackagesClient.BeginDelete method.
func (client *PackagesClient) BeginDelete(ctx context.Context, resourceGroupName string, testBaseAccountName string, packageName string, options *PackagesClientBeginDeleteOptions) (*runtime.Poller[PackagesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, testBaseAccountName, packageName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[PackagesClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[PackagesClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Deletes a Test Base Package.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-01-preview
func (client *PackagesClient) deleteOperation(ctx context.Context, resourceGroupName string, testBaseAccountName string, packageName string, options *PackagesClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "PackagesClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, testBaseAccountName, packageName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *PackagesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, testBaseAccountName string, packageName string, options *PackagesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TestBase/testBaseAccounts/{testBaseAccountName}/packages/{packageName}"
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
	if packageName == "" {
		return nil, errors.New("parameter packageName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{packageName}", url.PathEscape(packageName))
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

// Get - Gets a Test Base Package.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - testBaseAccountName - The resource name of the Test Base Account.
//   - packageName - The resource name of the Test Base Package.
//   - options - PackagesClientGetOptions contains the optional parameters for the PackagesClient.Get method.
func (client *PackagesClient) Get(ctx context.Context, resourceGroupName string, testBaseAccountName string, packageName string, options *PackagesClientGetOptions) (PackagesClientGetResponse, error) {
	var err error
	const operationName = "PackagesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, testBaseAccountName, packageName, options)
	if err != nil {
		return PackagesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PackagesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PackagesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *PackagesClient) getCreateRequest(ctx context.Context, resourceGroupName string, testBaseAccountName string, packageName string, options *PackagesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TestBase/testBaseAccounts/{testBaseAccountName}/packages/{packageName}"
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
	if packageName == "" {
		return nil, errors.New("parameter packageName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{packageName}", url.PathEscape(packageName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PackagesClient) getHandleResponse(resp *http.Response) (PackagesClientGetResponse, error) {
	result := PackagesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PackageResource); err != nil {
		return PackagesClientGetResponse{}, err
	}
	return result, nil
}

// GetDownloadURL - Gets the download URL of a package.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - testBaseAccountName - The resource name of the Test Base Account.
//   - packageName - The resource name of the Test Base Package.
//   - options - PackagesClientGetDownloadURLOptions contains the optional parameters for the PackagesClient.GetDownloadURL method.
func (client *PackagesClient) GetDownloadURL(ctx context.Context, resourceGroupName string, testBaseAccountName string, packageName string, options *PackagesClientGetDownloadURLOptions) (PackagesClientGetDownloadURLResponse, error) {
	var err error
	const operationName = "PackagesClient.GetDownloadURL"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getDownloadURLCreateRequest(ctx, resourceGroupName, testBaseAccountName, packageName, options)
	if err != nil {
		return PackagesClientGetDownloadURLResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PackagesClientGetDownloadURLResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PackagesClientGetDownloadURLResponse{}, err
	}
	resp, err := client.getDownloadURLHandleResponse(httpResp)
	return resp, err
}

// getDownloadURLCreateRequest creates the GetDownloadURL request.
func (client *PackagesClient) getDownloadURLCreateRequest(ctx context.Context, resourceGroupName string, testBaseAccountName string, packageName string, options *PackagesClientGetDownloadURLOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TestBase/testBaseAccounts/{testBaseAccountName}/packages/{packageName}/getDownloadUrl"
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
	if packageName == "" {
		return nil, errors.New("parameter packageName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{packageName}", url.PathEscape(packageName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getDownloadURLHandleResponse handles the GetDownloadURL response.
func (client *PackagesClient) getDownloadURLHandleResponse(resp *http.Response) (PackagesClientGetDownloadURLResponse, error) {
	result := PackagesClientGetDownloadURLResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DownloadURLResponse); err != nil {
		return PackagesClientGetDownloadURLResponse{}, err
	}
	return result, nil
}

// BeginHardDelete - Hard Delete a Test Base Package.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - testBaseAccountName - The resource name of the Test Base Account.
//   - packageName - The resource name of the Test Base Package.
//   - options - PackagesClientBeginHardDeleteOptions contains the optional parameters for the PackagesClient.BeginHardDelete
//     method.
func (client *PackagesClient) BeginHardDelete(ctx context.Context, resourceGroupName string, testBaseAccountName string, packageName string, options *PackagesClientBeginHardDeleteOptions) (*runtime.Poller[PackagesClientHardDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.hardDelete(ctx, resourceGroupName, testBaseAccountName, packageName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[PackagesClientHardDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[PackagesClientHardDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// HardDelete - Hard Delete a Test Base Package.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-01-preview
func (client *PackagesClient) hardDelete(ctx context.Context, resourceGroupName string, testBaseAccountName string, packageName string, options *PackagesClientBeginHardDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "PackagesClient.BeginHardDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.hardDeleteCreateRequest(ctx, resourceGroupName, testBaseAccountName, packageName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// hardDeleteCreateRequest creates the HardDelete request.
func (client *PackagesClient) hardDeleteCreateRequest(ctx context.Context, resourceGroupName string, testBaseAccountName string, packageName string, options *PackagesClientBeginHardDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TestBase/testBaseAccounts/{testBaseAccountName}/packages/{packageName}/hardDelete"
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
	if packageName == "" {
		return nil, errors.New("parameter packageName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{packageName}", url.PathEscape(packageName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// NewListByTestBaseAccountPager - Lists all the packages under a Test Base Account.
//
// Generated from API version 2023-11-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - testBaseAccountName - The resource name of the Test Base Account.
//   - options - PackagesClientListByTestBaseAccountOptions contains the optional parameters for the PackagesClient.NewListByTestBaseAccountPager
//     method.
func (client *PackagesClient) NewListByTestBaseAccountPager(resourceGroupName string, testBaseAccountName string, options *PackagesClientListByTestBaseAccountOptions) *runtime.Pager[PackagesClientListByTestBaseAccountResponse] {
	return runtime.NewPager(runtime.PagingHandler[PackagesClientListByTestBaseAccountResponse]{
		More: func(page PackagesClientListByTestBaseAccountResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *PackagesClientListByTestBaseAccountResponse) (PackagesClientListByTestBaseAccountResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "PackagesClient.NewListByTestBaseAccountPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByTestBaseAccountCreateRequest(ctx, resourceGroupName, testBaseAccountName, options)
			}, nil)
			if err != nil {
				return PackagesClientListByTestBaseAccountResponse{}, err
			}
			return client.listByTestBaseAccountHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByTestBaseAccountCreateRequest creates the ListByTestBaseAccount request.
func (client *PackagesClient) listByTestBaseAccountCreateRequest(ctx context.Context, resourceGroupName string, testBaseAccountName string, options *PackagesClientListByTestBaseAccountOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TestBase/testBaseAccounts/{testBaseAccountName}/packages"
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
	reqQP.Set("api-version", "2023-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByTestBaseAccountHandleResponse handles the ListByTestBaseAccount response.
func (client *PackagesClient) listByTestBaseAccountHandleResponse(resp *http.Response) (PackagesClientListByTestBaseAccountResponse, error) {
	result := PackagesClientListByTestBaseAccountResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PackageListResult); err != nil {
		return PackagesClientListByTestBaseAccountResponse{}, err
	}
	return result, nil
}

// RunTest - Trigger a test run on the package.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - testBaseAccountName - The resource name of the Test Base Account.
//   - packageName - The resource name of the Test Base Package.
//   - options - PackagesClientRunTestOptions contains the optional parameters for the PackagesClient.RunTest method.
func (client *PackagesClient) RunTest(ctx context.Context, resourceGroupName string, testBaseAccountName string, packageName string, options *PackagesClientRunTestOptions) (PackagesClientRunTestResponse, error) {
	var err error
	const operationName = "PackagesClient.RunTest"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.runTestCreateRequest(ctx, resourceGroupName, testBaseAccountName, packageName, options)
	if err != nil {
		return PackagesClientRunTestResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PackagesClientRunTestResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return PackagesClientRunTestResponse{}, err
	}
	resp, err := client.runTestHandleResponse(httpResp)
	return resp, err
}

// runTestCreateRequest creates the RunTest request.
func (client *PackagesClient) runTestCreateRequest(ctx context.Context, resourceGroupName string, testBaseAccountName string, packageName string, options *PackagesClientRunTestOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TestBase/testBaseAccounts/{testBaseAccountName}/packages/{packageName}/runTest"
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
	if packageName == "" {
		return nil, errors.New("parameter packageName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{packageName}", url.PathEscape(packageName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.Parameters != nil {
		if err := runtime.MarshalAsJSON(req, *options.Parameters); err != nil {
			return nil, err
		}
		return req, nil
	}
	return req, nil
}

// runTestHandleResponse handles the RunTest response.
func (client *PackagesClient) runTestHandleResponse(resp *http.Response) (PackagesClientRunTestResponse, error) {
	result := PackagesClientRunTestResponse{}
	if val := resp.Header.Get("Azure-AsyncOperation"); val != "" {
		result.AzureAsyncOperation = &val
	}
	if val := resp.Header.Get("Location"); val != "" {
		result.Location = &val
	}
	return result, nil
}

// BeginUpdate - Update an existing Test Base Package.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - testBaseAccountName - The resource name of the Test Base Account.
//   - packageName - The resource name of the Test Base Package.
//   - parameters - Parameters supplied to update a Test Base Package.
//   - options - PackagesClientBeginUpdateOptions contains the optional parameters for the PackagesClient.BeginUpdate method.
func (client *PackagesClient) BeginUpdate(ctx context.Context, resourceGroupName string, testBaseAccountName string, packageName string, parameters PackageUpdateParameters, options *PackagesClientBeginUpdateOptions) (*runtime.Poller[PackagesClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, testBaseAccountName, packageName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[PackagesClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[PackagesClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Update - Update an existing Test Base Package.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-01-preview
func (client *PackagesClient) update(ctx context.Context, resourceGroupName string, testBaseAccountName string, packageName string, parameters PackageUpdateParameters, options *PackagesClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "PackagesClient.BeginUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, testBaseAccountName, packageName, parameters, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// updateCreateRequest creates the Update request.
func (client *PackagesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, testBaseAccountName string, packageName string, parameters PackageUpdateParameters, options *PackagesClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TestBase/testBaseAccounts/{testBaseAccountName}/packages/{packageName}"
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
	if packageName == "" {
		return nil, errors.New("parameter packageName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{packageName}", url.PathEscape(packageName))
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
