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

// ManagedInstanceLongTermRetentionPoliciesClient contains the methods for the ManagedInstanceLongTermRetentionPolicies group.
// Don't use this type directly, use NewManagedInstanceLongTermRetentionPoliciesClient() instead.
type ManagedInstanceLongTermRetentionPoliciesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewManagedInstanceLongTermRetentionPoliciesClient creates a new instance of ManagedInstanceLongTermRetentionPoliciesClient with the specified values.
//   - subscriptionID - The subscription ID that identifies an Azure subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewManagedInstanceLongTermRetentionPoliciesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ManagedInstanceLongTermRetentionPoliciesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ManagedInstanceLongTermRetentionPoliciesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Sets a managed database's long term retention policy.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-11-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - managedInstanceName - The name of the managed instance.
//   - databaseName - The name of the database.
//   - policyName - The policy name. Should always be Default.
//   - parameters - The long term retention policy info.
//   - options - ManagedInstanceLongTermRetentionPoliciesClientBeginCreateOrUpdateOptions contains the optional parameters for
//     the ManagedInstanceLongTermRetentionPoliciesClient.BeginCreateOrUpdate method.
func (client *ManagedInstanceLongTermRetentionPoliciesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, policyName ManagedInstanceLongTermRetentionPolicyName, parameters ManagedInstanceLongTermRetentionPolicy, options *ManagedInstanceLongTermRetentionPoliciesClientBeginCreateOrUpdateOptions) (*runtime.Poller[ManagedInstanceLongTermRetentionPoliciesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, managedInstanceName, databaseName, policyName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ManagedInstanceLongTermRetentionPoliciesClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ManagedInstanceLongTermRetentionPoliciesClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Sets a managed database's long term retention policy.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-11-01-preview
func (client *ManagedInstanceLongTermRetentionPoliciesClient) createOrUpdate(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, policyName ManagedInstanceLongTermRetentionPolicyName, parameters ManagedInstanceLongTermRetentionPolicy, options *ManagedInstanceLongTermRetentionPoliciesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "ManagedInstanceLongTermRetentionPoliciesClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
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
func (client *ManagedInstanceLongTermRetentionPoliciesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, policyName ManagedInstanceLongTermRetentionPolicyName, parameters ManagedInstanceLongTermRetentionPolicy, options *ManagedInstanceLongTermRetentionPoliciesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/databases/{databaseName}/backupLongTermRetentionPolicies/{policyName}"
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
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// Get - Gets a managed database's long term retention policy.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-11-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - managedInstanceName - The name of the managed instance.
//   - databaseName - The name of the database.
//   - policyName - The policy name. Should always be Default.
//   - options - ManagedInstanceLongTermRetentionPoliciesClientGetOptions contains the optional parameters for the ManagedInstanceLongTermRetentionPoliciesClient.Get
//     method.
func (client *ManagedInstanceLongTermRetentionPoliciesClient) Get(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, policyName ManagedInstanceLongTermRetentionPolicyName, options *ManagedInstanceLongTermRetentionPoliciesClientGetOptions) (ManagedInstanceLongTermRetentionPoliciesClientGetResponse, error) {
	var err error
	const operationName = "ManagedInstanceLongTermRetentionPoliciesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, managedInstanceName, databaseName, policyName, options)
	if err != nil {
		return ManagedInstanceLongTermRetentionPoliciesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ManagedInstanceLongTermRetentionPoliciesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ManagedInstanceLongTermRetentionPoliciesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ManagedInstanceLongTermRetentionPoliciesClient) getCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, policyName ManagedInstanceLongTermRetentionPolicyName, options *ManagedInstanceLongTermRetentionPoliciesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/databases/{databaseName}/backupLongTermRetentionPolicies/{policyName}"
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
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ManagedInstanceLongTermRetentionPoliciesClient) getHandleResponse(resp *http.Response) (ManagedInstanceLongTermRetentionPoliciesClientGetResponse, error) {
	result := ManagedInstanceLongTermRetentionPoliciesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagedInstanceLongTermRetentionPolicy); err != nil {
		return ManagedInstanceLongTermRetentionPoliciesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByDatabasePager - Gets a database's long term retention policy.
//
// Generated from API version 2020-11-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - managedInstanceName - The name of the managed instance.
//   - databaseName - The name of the database.
//   - options - ManagedInstanceLongTermRetentionPoliciesClientListByDatabaseOptions contains the optional parameters for the
//     ManagedInstanceLongTermRetentionPoliciesClient.NewListByDatabasePager method.
func (client *ManagedInstanceLongTermRetentionPoliciesClient) NewListByDatabasePager(resourceGroupName string, managedInstanceName string, databaseName string, options *ManagedInstanceLongTermRetentionPoliciesClientListByDatabaseOptions) *runtime.Pager[ManagedInstanceLongTermRetentionPoliciesClientListByDatabaseResponse] {
	return runtime.NewPager(runtime.PagingHandler[ManagedInstanceLongTermRetentionPoliciesClientListByDatabaseResponse]{
		More: func(page ManagedInstanceLongTermRetentionPoliciesClientListByDatabaseResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ManagedInstanceLongTermRetentionPoliciesClientListByDatabaseResponse) (ManagedInstanceLongTermRetentionPoliciesClientListByDatabaseResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ManagedInstanceLongTermRetentionPoliciesClient.NewListByDatabasePager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByDatabaseCreateRequest(ctx, resourceGroupName, managedInstanceName, databaseName, options)
			}, nil)
			if err != nil {
				return ManagedInstanceLongTermRetentionPoliciesClientListByDatabaseResponse{}, err
			}
			return client.listByDatabaseHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByDatabaseCreateRequest creates the ListByDatabase request.
func (client *ManagedInstanceLongTermRetentionPoliciesClient) listByDatabaseCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, options *ManagedInstanceLongTermRetentionPoliciesClientListByDatabaseOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/databases/{databaseName}/backupLongTermRetentionPolicies"
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
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByDatabaseHandleResponse handles the ListByDatabase response.
func (client *ManagedInstanceLongTermRetentionPoliciesClient) listByDatabaseHandleResponse(resp *http.Response) (ManagedInstanceLongTermRetentionPoliciesClientListByDatabaseResponse, error) {
	result := ManagedInstanceLongTermRetentionPoliciesClientListByDatabaseResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagedInstanceLongTermRetentionPolicyListResult); err != nil {
		return ManagedInstanceLongTermRetentionPoliciesClientListByDatabaseResponse{}, err
	}
	return result, nil
}
