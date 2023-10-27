//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armautomation

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
//   - subscriptionID - Gets subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID
//     forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewPackagesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*PackagesClient, error) {
	cl, err := arm.NewClient(moduleName+".PackagesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &PackagesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewListByRuntimeEnvironmentPager - Retrieve the a list of Packages
//
// Generated from API version 2023-05-15-preview
//   - resourceGroupName - Name of an Azure Resource group.
//   - automationAccountName - The name of the automation account.
//   - runtimeEnvironmentName - The name of the Runtime Environment.
//   - options - PackagesClientListByRuntimeEnvironmentOptions contains the optional parameters for the PackagesClient.NewListByRuntimeEnvironmentPager
//     method.
func (client *PackagesClient) NewListByRuntimeEnvironmentPager(resourceGroupName string, automationAccountName string, runtimeEnvironmentName string, options *PackagesClientListByRuntimeEnvironmentOptions) *runtime.Pager[PackagesClientListByRuntimeEnvironmentResponse] {
	return runtime.NewPager(runtime.PagingHandler[PackagesClientListByRuntimeEnvironmentResponse]{
		More: func(page PackagesClientListByRuntimeEnvironmentResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *PackagesClientListByRuntimeEnvironmentResponse) (PackagesClientListByRuntimeEnvironmentResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByRuntimeEnvironmentCreateRequest(ctx, resourceGroupName, automationAccountName, runtimeEnvironmentName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return PackagesClientListByRuntimeEnvironmentResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return PackagesClientListByRuntimeEnvironmentResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return PackagesClientListByRuntimeEnvironmentResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByRuntimeEnvironmentHandleResponse(resp)
		},
	})
}

// listByRuntimeEnvironmentCreateRequest creates the ListByRuntimeEnvironment request.
func (client *PackagesClient) listByRuntimeEnvironmentCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, runtimeEnvironmentName string, options *PackagesClientListByRuntimeEnvironmentOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/runtimeEnvironments/{runtimeEnvironmentName}/packages"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if runtimeEnvironmentName == "" {
		return nil, errors.New("parameter runtimeEnvironmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{runtimeEnvironmentName}", url.PathEscape(runtimeEnvironmentName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByRuntimeEnvironmentHandleResponse handles the ListByRuntimeEnvironment response.
func (client *PackagesClient) listByRuntimeEnvironmentHandleResponse(resp *http.Response) (PackagesClientListByRuntimeEnvironmentResponse, error) {
	result := PackagesClientListByRuntimeEnvironmentResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PackageListResult); err != nil {
		return PackagesClientListByRuntimeEnvironmentResponse{}, err
	}
	return result, nil
}
