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

// BackupPoliciesClient contains the methods for the BackupPolicies group.
// Don't use this type directly, use NewBackupPoliciesClient() instead.
type BackupPoliciesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewBackupPoliciesClient creates a new instance of BackupPoliciesClient with the specified values.
//   - subscriptionID - The subscription Id.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewBackupPoliciesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*BackupPoliciesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &BackupPoliciesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewListPager - Lists of backup policies associated with Recovery Services Vault. API provides pagination parameters to
// fetch scoped results.
//
// Generated from API version 2025-02-01
//   - vaultName - The name of the recovery services vault.
//   - resourceGroupName - The name of the resource group where the recovery services vault is present.
//   - options - BackupPoliciesClientListOptions contains the optional parameters for the BackupPoliciesClient.NewListPager method.
func (client *BackupPoliciesClient) NewListPager(vaultName string, resourceGroupName string, options *BackupPoliciesClientListOptions) *runtime.Pager[BackupPoliciesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[BackupPoliciesClientListResponse]{
		More: func(page BackupPoliciesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *BackupPoliciesClientListResponse) (BackupPoliciesClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "BackupPoliciesClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, vaultName, resourceGroupName, options)
			}, nil)
			if err != nil {
				return BackupPoliciesClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *BackupPoliciesClient) listCreateRequest(ctx context.Context, vaultName string, resourceGroupName string, options *BackupPoliciesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupPolicies"
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	reqQP.Set("api-version", "2025-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *BackupPoliciesClient) listHandleResponse(resp *http.Response) (BackupPoliciesClientListResponse, error) {
	result := BackupPoliciesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProtectionPolicyResourceList); err != nil {
		return BackupPoliciesClientListResponse{}, err
	}
	return result, nil
}
