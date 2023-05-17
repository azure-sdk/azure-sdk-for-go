//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

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

// ManagedInstanceAdministratorsClient contains the methods for the ManagedInstanceAdministrators group.
// Don't use this type directly, use NewManagedInstanceAdministratorsClient() instead.
type ManagedInstanceAdministratorsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewManagedInstanceAdministratorsClient creates a new instance of ManagedInstanceAdministratorsClient with the specified values.
//   - subscriptionID - The subscription ID that identifies an Azure subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewManagedInstanceAdministratorsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ManagedInstanceAdministratorsClient, error) {
	cl, err := arm.NewClient(moduleName+".ManagedInstanceAdministratorsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ManagedInstanceAdministratorsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates a managed instance administrator.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - managedInstanceName - The name of the managed instance.
//   - parameters - The requested administrator parameters.
//   - options - ManagedInstanceAdministratorsClientBeginCreateOrUpdateOptions contains the optional parameters for the ManagedInstanceAdministratorsClient.BeginCreateOrUpdate
//     method.
func (client *ManagedInstanceAdministratorsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, managedInstanceName string, administratorName AdministratorName, parameters ManagedInstanceAdministrator, options *ManagedInstanceAdministratorsClientBeginCreateOrUpdateOptions) (*runtime.Poller[ManagedInstanceAdministratorsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, managedInstanceName, administratorName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[ManagedInstanceAdministratorsClientCreateOrUpdateResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[ManagedInstanceAdministratorsClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Creates or updates a managed instance administrator.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01-preview
func (client *ManagedInstanceAdministratorsClient) createOrUpdate(ctx context.Context, resourceGroupName string, managedInstanceName string, administratorName AdministratorName, parameters ManagedInstanceAdministrator, options *ManagedInstanceAdministratorsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, managedInstanceName, administratorName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ManagedInstanceAdministratorsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, administratorName AdministratorName, parameters ManagedInstanceAdministrator, options *ManagedInstanceAdministratorsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/administrators/{administratorName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if administratorName == "" {
		return nil, errors.New("parameter administratorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{administratorName}", url.PathEscape(string(administratorName)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Deletes a managed instance administrator.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - managedInstanceName - The name of the managed instance.
//   - options - ManagedInstanceAdministratorsClientBeginDeleteOptions contains the optional parameters for the ManagedInstanceAdministratorsClient.BeginDelete
//     method.
func (client *ManagedInstanceAdministratorsClient) BeginDelete(ctx context.Context, resourceGroupName string, managedInstanceName string, administratorName AdministratorName, options *ManagedInstanceAdministratorsClientBeginDeleteOptions) (*runtime.Poller[ManagedInstanceAdministratorsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, managedInstanceName, administratorName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[ManagedInstanceAdministratorsClientDeleteResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[ManagedInstanceAdministratorsClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Deletes a managed instance administrator.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01-preview
func (client *ManagedInstanceAdministratorsClient) deleteOperation(ctx context.Context, resourceGroupName string, managedInstanceName string, administratorName AdministratorName, options *ManagedInstanceAdministratorsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, managedInstanceName, administratorName, options)
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

// deleteCreateRequest creates the Delete request.
func (client *ManagedInstanceAdministratorsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, administratorName AdministratorName, options *ManagedInstanceAdministratorsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/administrators/{administratorName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if administratorName == "" {
		return nil, errors.New("parameter administratorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{administratorName}", url.PathEscape(string(administratorName)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Gets a managed instance administrator.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - managedInstanceName - The name of the managed instance.
//   - options - ManagedInstanceAdministratorsClientGetOptions contains the optional parameters for the ManagedInstanceAdministratorsClient.Get
//     method.
func (client *ManagedInstanceAdministratorsClient) Get(ctx context.Context, resourceGroupName string, managedInstanceName string, administratorName AdministratorName, options *ManagedInstanceAdministratorsClientGetOptions) (ManagedInstanceAdministratorsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, managedInstanceName, administratorName, options)
	if err != nil {
		return ManagedInstanceAdministratorsClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ManagedInstanceAdministratorsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ManagedInstanceAdministratorsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ManagedInstanceAdministratorsClient) getCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, administratorName AdministratorName, options *ManagedInstanceAdministratorsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/administrators/{administratorName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if administratorName == "" {
		return nil, errors.New("parameter administratorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{administratorName}", url.PathEscape(string(administratorName)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ManagedInstanceAdministratorsClient) getHandleResponse(resp *http.Response) (ManagedInstanceAdministratorsClientGetResponse, error) {
	result := ManagedInstanceAdministratorsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagedInstanceAdministrator); err != nil {
		return ManagedInstanceAdministratorsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByInstancePager - Gets a list of managed instance administrators.
//
// Generated from API version 2022-11-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - managedInstanceName - The name of the managed instance.
//   - options - ManagedInstanceAdministratorsClientListByInstanceOptions contains the optional parameters for the ManagedInstanceAdministratorsClient.NewListByInstancePager
//     method.
func (client *ManagedInstanceAdministratorsClient) NewListByInstancePager(resourceGroupName string, managedInstanceName string, options *ManagedInstanceAdministratorsClientListByInstanceOptions) *runtime.Pager[ManagedInstanceAdministratorsClientListByInstanceResponse] {
	return runtime.NewPager(runtime.PagingHandler[ManagedInstanceAdministratorsClientListByInstanceResponse]{
		More: func(page ManagedInstanceAdministratorsClientListByInstanceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ManagedInstanceAdministratorsClientListByInstanceResponse) (ManagedInstanceAdministratorsClientListByInstanceResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByInstanceCreateRequest(ctx, resourceGroupName, managedInstanceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ManagedInstanceAdministratorsClientListByInstanceResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ManagedInstanceAdministratorsClientListByInstanceResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ManagedInstanceAdministratorsClientListByInstanceResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByInstanceHandleResponse(resp)
		},
	})
}

// listByInstanceCreateRequest creates the ListByInstance request.
func (client *ManagedInstanceAdministratorsClient) listByInstanceCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, options *ManagedInstanceAdministratorsClientListByInstanceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/administrators"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByInstanceHandleResponse handles the ListByInstance response.
func (client *ManagedInstanceAdministratorsClient) listByInstanceHandleResponse(resp *http.Response) (ManagedInstanceAdministratorsClientListByInstanceResponse, error) {
	result := ManagedInstanceAdministratorsClientListByInstanceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagedInstanceAdministratorListResult); err != nil {
		return ManagedInstanceAdministratorsClientListByInstanceResponse{}, err
	}
	return result, nil
}
