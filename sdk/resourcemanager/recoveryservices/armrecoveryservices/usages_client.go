//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

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

// UsagesClient contains the methods for the Usages group.
// Don't use this type directly, use NewUsagesClient() instead.
type UsagesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewUsagesClient creates a new instance of UsagesClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewUsagesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*UsagesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &UsagesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewListByVaultsPager - Fetches the usages of the vault.
//
// Generated from API version 2024-10-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - vaultName - The name of the recovery services vault.
//   - options - UsagesClientListByVaultsOptions contains the optional parameters for the UsagesClient.NewListByVaultsPager method.
func (client *UsagesClient) NewListByVaultsPager(resourceGroupName string, vaultName string, options *UsagesClientListByVaultsOptions) *runtime.Pager[UsagesClientListByVaultsResponse] {
	return runtime.NewPager(runtime.PagingHandler[UsagesClientListByVaultsResponse]{
		More: func(page UsagesClientListByVaultsResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *UsagesClientListByVaultsResponse) (UsagesClientListByVaultsResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "UsagesClient.NewListByVaultsPager")
			req, err := client.listByVaultsCreateRequest(ctx, resourceGroupName, vaultName, options)
			if err != nil {
				return UsagesClientListByVaultsResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return UsagesClientListByVaultsResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return UsagesClientListByVaultsResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByVaultsHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByVaultsCreateRequest creates the ListByVaults request.
func (client *UsagesClient) listByVaultsCreateRequest(ctx context.Context, resourceGroupName string, vaultName string, options *UsagesClientListByVaultsOptions) (*policy.Request, error) {
	urlPath := "/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/usages"
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
	reqQP.Set("api-version", "2024-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByVaultsHandleResponse handles the ListByVaults response.
func (client *UsagesClient) listByVaultsHandleResponse(resp *http.Response) (UsagesClientListByVaultsResponse, error) {
	result := UsagesClientListByVaultsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VaultUsageList); err != nil {
		return UsagesClientListByVaultsResponse{}, err
	}
	return result, nil
}
