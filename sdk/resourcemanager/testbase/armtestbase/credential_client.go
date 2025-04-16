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

// CredentialClient contains the methods for the Credential group.
// Don't use this type directly, use NewCredentialClient() instead.
type CredentialClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewCredentialClient creates a new instance of CredentialClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewCredentialClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*CredentialClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &CredentialClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Gets a test base credential Resource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - testBaseAccountName - The resource name of the Test Base Account.
//   - credentialName - The credential resource name.
//   - options - CredentialClientGetOptions contains the optional parameters for the CredentialClient.Get method.
func (client *CredentialClient) Get(ctx context.Context, resourceGroupName string, testBaseAccountName string, credentialName string, options *CredentialClientGetOptions) (CredentialClientGetResponse, error) {
	var err error
	const operationName = "CredentialClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, testBaseAccountName, credentialName, options)
	if err != nil {
		return CredentialClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return CredentialClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return CredentialClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *CredentialClient) getCreateRequest(ctx context.Context, resourceGroupName string, testBaseAccountName string, credentialName string, _ *CredentialClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TestBase/testBaseAccounts/{testBaseAccountName}/credentials/{credentialName}"
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
	if credentialName == "" {
		return nil, errors.New("parameter credentialName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{credentialName}", url.PathEscape(credentialName))
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
func (client *CredentialClient) getHandleResponse(resp *http.Response) (CredentialClientGetResponse, error) {
	result := CredentialClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CredentialResource); err != nil {
		return CredentialClientGetResponse{}, err
	}
	return result, nil
}

// NewListByTestBaseAccountPager - Lists all the Credentials under a Test Base Account.
//
// Generated from API version 2023-11-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - testBaseAccountName - The resource name of the Test Base Account.
//   - options - CredentialClientListByTestBaseAccountOptions contains the optional parameters for the CredentialClient.NewListByTestBaseAccountPager
//     method.
func (client *CredentialClient) NewListByTestBaseAccountPager(resourceGroupName string, testBaseAccountName string, options *CredentialClientListByTestBaseAccountOptions) *runtime.Pager[CredentialClientListByTestBaseAccountResponse] {
	return runtime.NewPager(runtime.PagingHandler[CredentialClientListByTestBaseAccountResponse]{
		More: func(page CredentialClientListByTestBaseAccountResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *CredentialClientListByTestBaseAccountResponse) (CredentialClientListByTestBaseAccountResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "CredentialClient.NewListByTestBaseAccountPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByTestBaseAccountCreateRequest(ctx, resourceGroupName, testBaseAccountName, options)
			}, nil)
			if err != nil {
				return CredentialClientListByTestBaseAccountResponse{}, err
			}
			return client.listByTestBaseAccountHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByTestBaseAccountCreateRequest creates the ListByTestBaseAccount request.
func (client *CredentialClient) listByTestBaseAccountCreateRequest(ctx context.Context, resourceGroupName string, testBaseAccountName string, _ *CredentialClientListByTestBaseAccountOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TestBase/testBaseAccounts/{testBaseAccountName}/credentials"
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
func (client *CredentialClient) listByTestBaseAccountHandleResponse(resp *http.Response) (CredentialClientListByTestBaseAccountResponse, error) {
	result := CredentialClientListByTestBaseAccountResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CredentialListResult); err != nil {
		return CredentialClientListByTestBaseAccountResponse{}, err
	}
	return result, nil
}
