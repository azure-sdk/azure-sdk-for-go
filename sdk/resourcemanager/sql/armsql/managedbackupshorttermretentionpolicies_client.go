//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsql

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

// ManagedBackupShortTermRetentionPoliciesClient contains the methods for the ManagedBackupShortTermRetentionPolicies group.
// Don't use this type directly, use NewManagedBackupShortTermRetentionPoliciesClient() instead.
type ManagedBackupShortTermRetentionPoliciesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewManagedBackupShortTermRetentionPoliciesClient creates a new instance of ManagedBackupShortTermRetentionPoliciesClient with the specified values.
//   - subscriptionID - The subscription ID that identifies an Azure subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewManagedBackupShortTermRetentionPoliciesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ManagedBackupShortTermRetentionPoliciesClient, error) {
	cl, err := arm.NewClient(moduleName+".ManagedBackupShortTermRetentionPoliciesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ManagedBackupShortTermRetentionPoliciesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Updates a managed database's short term retention policy.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-02-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - managedInstanceName - The name of the managed instance.
//   - databaseName - The name of the database.
//   - policyName - The policy name. Should always be "default".
//   - parameters - The short term retention policy info.
//   - options - ManagedBackupShortTermRetentionPoliciesClientBeginCreateOrUpdateOptions contains the optional parameters for
//     the ManagedBackupShortTermRetentionPoliciesClient.BeginCreateOrUpdate method.
func (client *ManagedBackupShortTermRetentionPoliciesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, policyName ManagedShortTermRetentionPolicyName, parameters ManagedBackupShortTermRetentionPolicy, options *ManagedBackupShortTermRetentionPoliciesClientBeginCreateOrUpdateOptions) (*runtime.Poller[ManagedBackupShortTermRetentionPoliciesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, managedInstanceName, databaseName, policyName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[ManagedBackupShortTermRetentionPoliciesClientCreateOrUpdateResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[ManagedBackupShortTermRetentionPoliciesClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Updates a managed database's short term retention policy.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-02-01-preview
func (client *ManagedBackupShortTermRetentionPoliciesClient) createOrUpdate(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, policyName ManagedShortTermRetentionPolicyName, parameters ManagedBackupShortTermRetentionPolicy, options *ManagedBackupShortTermRetentionPoliciesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, managedInstanceName, databaseName, policyName, parameters, options)
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

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ManagedBackupShortTermRetentionPoliciesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, policyName ManagedShortTermRetentionPolicyName, parameters ManagedBackupShortTermRetentionPolicy, options *ManagedBackupShortTermRetentionPoliciesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/databases/{databaseName}/backupShortTermRetentionPolicies/{policyName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if policyName == "" {
		return nil, errors.New("parameter policyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{policyName}", url.PathEscape(string(policyName)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// Get - Gets a managed database's short term retention policy.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-02-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - managedInstanceName - The name of the managed instance.
//   - databaseName - The name of the database.
//   - policyName - The policy name.
//   - options - ManagedBackupShortTermRetentionPoliciesClientGetOptions contains the optional parameters for the ManagedBackupShortTermRetentionPoliciesClient.Get
//     method.
func (client *ManagedBackupShortTermRetentionPoliciesClient) Get(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, policyName ManagedShortTermRetentionPolicyName, options *ManagedBackupShortTermRetentionPoliciesClientGetOptions) (ManagedBackupShortTermRetentionPoliciesClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, managedInstanceName, databaseName, policyName, options)
	if err != nil {
		return ManagedBackupShortTermRetentionPoliciesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ManagedBackupShortTermRetentionPoliciesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ManagedBackupShortTermRetentionPoliciesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ManagedBackupShortTermRetentionPoliciesClient) getCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, policyName ManagedShortTermRetentionPolicyName, options *ManagedBackupShortTermRetentionPoliciesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/databases/{databaseName}/backupShortTermRetentionPolicies/{policyName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if policyName == "" {
		return nil, errors.New("parameter policyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{policyName}", url.PathEscape(string(policyName)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ManagedBackupShortTermRetentionPoliciesClient) getHandleResponse(resp *http.Response) (ManagedBackupShortTermRetentionPoliciesClientGetResponse, error) {
	result := ManagedBackupShortTermRetentionPoliciesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagedBackupShortTermRetentionPolicy); err != nil {
		return ManagedBackupShortTermRetentionPoliciesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByDatabasePager - Gets a managed database's short term retention policy list.
//
// Generated from API version 2023-02-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - managedInstanceName - The name of the managed instance.
//   - databaseName - The name of the database.
//   - options - ManagedBackupShortTermRetentionPoliciesClientListByDatabaseOptions contains the optional parameters for the ManagedBackupShortTermRetentionPoliciesClient.NewListByDatabasePager
//     method.
func (client *ManagedBackupShortTermRetentionPoliciesClient) NewListByDatabasePager(resourceGroupName string, managedInstanceName string, databaseName string, options *ManagedBackupShortTermRetentionPoliciesClientListByDatabaseOptions) *runtime.Pager[ManagedBackupShortTermRetentionPoliciesClientListByDatabaseResponse] {
	return runtime.NewPager(runtime.PagingHandler[ManagedBackupShortTermRetentionPoliciesClientListByDatabaseResponse]{
		More: func(page ManagedBackupShortTermRetentionPoliciesClientListByDatabaseResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ManagedBackupShortTermRetentionPoliciesClientListByDatabaseResponse) (ManagedBackupShortTermRetentionPoliciesClientListByDatabaseResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByDatabaseCreateRequest(ctx, resourceGroupName, managedInstanceName, databaseName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ManagedBackupShortTermRetentionPoliciesClientListByDatabaseResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ManagedBackupShortTermRetentionPoliciesClientListByDatabaseResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ManagedBackupShortTermRetentionPoliciesClientListByDatabaseResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByDatabaseHandleResponse(resp)
		},
	})
}

// listByDatabaseCreateRequest creates the ListByDatabase request.
func (client *ManagedBackupShortTermRetentionPoliciesClient) listByDatabaseCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, options *ManagedBackupShortTermRetentionPoliciesClientListByDatabaseOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/databases/{databaseName}/backupShortTermRetentionPolicies"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByDatabaseHandleResponse handles the ListByDatabase response.
func (client *ManagedBackupShortTermRetentionPoliciesClient) listByDatabaseHandleResponse(resp *http.Response) (ManagedBackupShortTermRetentionPoliciesClientListByDatabaseResponse, error) {
	result := ManagedBackupShortTermRetentionPoliciesClientListByDatabaseResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagedBackupShortTermRetentionPolicyListResult); err != nil {
		return ManagedBackupShortTermRetentionPoliciesClientListByDatabaseResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Updates a managed database's short term retention policy.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-02-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - managedInstanceName - The name of the managed instance.
//   - databaseName - The name of the database.
//   - policyName - The policy name. Should always be "default".
//   - parameters - The short term retention policy info.
//   - options - ManagedBackupShortTermRetentionPoliciesClientBeginUpdateOptions contains the optional parameters for the ManagedBackupShortTermRetentionPoliciesClient.BeginUpdate
//     method.
func (client *ManagedBackupShortTermRetentionPoliciesClient) BeginUpdate(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, policyName ManagedShortTermRetentionPolicyName, parameters ManagedBackupShortTermRetentionPolicy, options *ManagedBackupShortTermRetentionPoliciesClientBeginUpdateOptions) (*runtime.Poller[ManagedBackupShortTermRetentionPoliciesClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, managedInstanceName, databaseName, policyName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[ManagedBackupShortTermRetentionPoliciesClientUpdateResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[ManagedBackupShortTermRetentionPoliciesClientUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Update - Updates a managed database's short term retention policy.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-02-01-preview
func (client *ManagedBackupShortTermRetentionPoliciesClient) update(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, policyName ManagedShortTermRetentionPolicyName, parameters ManagedBackupShortTermRetentionPolicy, options *ManagedBackupShortTermRetentionPoliciesClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	req, err := client.updateCreateRequest(ctx, resourceGroupName, managedInstanceName, databaseName, policyName, parameters, options)
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

// updateCreateRequest creates the Update request.
func (client *ManagedBackupShortTermRetentionPoliciesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, policyName ManagedShortTermRetentionPolicyName, parameters ManagedBackupShortTermRetentionPolicy, options *ManagedBackupShortTermRetentionPoliciesClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/databases/{databaseName}/backupShortTermRetentionPolicies/{policyName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if policyName == "" {
		return nil, errors.New("parameter policyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{policyName}", url.PathEscape(string(policyName)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}
