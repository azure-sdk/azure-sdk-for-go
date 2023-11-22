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

// ManagedInstanceAzureADOnlyAuthenticationsClient contains the methods for the ManagedInstanceAzureADOnlyAuthentications group.
// Don't use this type directly, use NewManagedInstanceAzureADOnlyAuthenticationsClient() instead.
type ManagedInstanceAzureADOnlyAuthenticationsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewManagedInstanceAzureADOnlyAuthenticationsClient creates a new instance of ManagedInstanceAzureADOnlyAuthenticationsClient with the specified values.
//   - subscriptionID - The subscription ID that identifies an Azure subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewManagedInstanceAzureADOnlyAuthenticationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ManagedInstanceAzureADOnlyAuthenticationsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ManagedInstanceAzureADOnlyAuthenticationsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Sets Server Active Directory only authentication property or updates an existing server Active Directory
// only authentication property.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-08-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - managedInstanceName - The name of the managed instance.
//   - authenticationName - The name of server azure active directory only authentication.
//   - parameters - The required parameters for creating or updating an Active Directory only authentication property.
//   - options - ManagedInstanceAzureADOnlyAuthenticationsClientBeginCreateOrUpdateOptions contains the optional parameters for
//     the ManagedInstanceAzureADOnlyAuthenticationsClient.BeginCreateOrUpdate method.
func (client *ManagedInstanceAzureADOnlyAuthenticationsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, managedInstanceName string, authenticationName AuthenticationName, parameters ManagedInstanceAzureADOnlyAuthentication, options *ManagedInstanceAzureADOnlyAuthenticationsClientBeginCreateOrUpdateOptions) (*runtime.Poller[ManagedInstanceAzureADOnlyAuthenticationsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, managedInstanceName, authenticationName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ManagedInstanceAzureADOnlyAuthenticationsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ManagedInstanceAzureADOnlyAuthenticationsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Sets Server Active Directory only authentication property or updates an existing server Active Directory
// only authentication property.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-08-01-preview
func (client *ManagedInstanceAzureADOnlyAuthenticationsClient) createOrUpdate(ctx context.Context, resourceGroupName string, managedInstanceName string, authenticationName AuthenticationName, parameters ManagedInstanceAzureADOnlyAuthentication, options *ManagedInstanceAzureADOnlyAuthenticationsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "ManagedInstanceAzureADOnlyAuthenticationsClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, managedInstanceName, authenticationName, parameters, options)
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

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ManagedInstanceAzureADOnlyAuthenticationsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, authenticationName AuthenticationName, parameters ManagedInstanceAzureADOnlyAuthentication, options *ManagedInstanceAzureADOnlyAuthenticationsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/azureADOnlyAuthentications/{authenticationName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if authenticationName == "" {
		return nil, errors.New("parameter authenticationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{authenticationName}", url.PathEscape(string(authenticationName)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes an existing server Active Directory only authentication property.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-08-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - managedInstanceName - The name of the managed instance.
//   - authenticationName - The name of server azure active directory only authentication.
//   - options - ManagedInstanceAzureADOnlyAuthenticationsClientBeginDeleteOptions contains the optional parameters for the ManagedInstanceAzureADOnlyAuthenticationsClient.BeginDelete
//     method.
func (client *ManagedInstanceAzureADOnlyAuthenticationsClient) BeginDelete(ctx context.Context, resourceGroupName string, managedInstanceName string, authenticationName AuthenticationName, options *ManagedInstanceAzureADOnlyAuthenticationsClientBeginDeleteOptions) (*runtime.Poller[ManagedInstanceAzureADOnlyAuthenticationsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, managedInstanceName, authenticationName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ManagedInstanceAzureADOnlyAuthenticationsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ManagedInstanceAzureADOnlyAuthenticationsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Deletes an existing server Active Directory only authentication property.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-08-01-preview
func (client *ManagedInstanceAzureADOnlyAuthenticationsClient) deleteOperation(ctx context.Context, resourceGroupName string, managedInstanceName string, authenticationName AuthenticationName, options *ManagedInstanceAzureADOnlyAuthenticationsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "ManagedInstanceAzureADOnlyAuthenticationsClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, managedInstanceName, authenticationName, options)
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
func (client *ManagedInstanceAzureADOnlyAuthenticationsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, authenticationName AuthenticationName, options *ManagedInstanceAzureADOnlyAuthenticationsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/azureADOnlyAuthentications/{authenticationName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if authenticationName == "" {
		return nil, errors.New("parameter authenticationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{authenticationName}", url.PathEscape(string(authenticationName)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Gets a specific Azure Active Directory only authentication property.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-08-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - managedInstanceName - The name of the managed instance.
//   - authenticationName - The name of server azure active directory only authentication.
//   - options - ManagedInstanceAzureADOnlyAuthenticationsClientGetOptions contains the optional parameters for the ManagedInstanceAzureADOnlyAuthenticationsClient.Get
//     method.
func (client *ManagedInstanceAzureADOnlyAuthenticationsClient) Get(ctx context.Context, resourceGroupName string, managedInstanceName string, authenticationName AuthenticationName, options *ManagedInstanceAzureADOnlyAuthenticationsClientGetOptions) (ManagedInstanceAzureADOnlyAuthenticationsClientGetResponse, error) {
	var err error
	const operationName = "ManagedInstanceAzureADOnlyAuthenticationsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, managedInstanceName, authenticationName, options)
	if err != nil {
		return ManagedInstanceAzureADOnlyAuthenticationsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ManagedInstanceAzureADOnlyAuthenticationsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ManagedInstanceAzureADOnlyAuthenticationsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ManagedInstanceAzureADOnlyAuthenticationsClient) getCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, authenticationName AuthenticationName, options *ManagedInstanceAzureADOnlyAuthenticationsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/azureADOnlyAuthentications/{authenticationName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if authenticationName == "" {
		return nil, errors.New("parameter authenticationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{authenticationName}", url.PathEscape(string(authenticationName)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ManagedInstanceAzureADOnlyAuthenticationsClient) getHandleResponse(resp *http.Response) (ManagedInstanceAzureADOnlyAuthenticationsClientGetResponse, error) {
	result := ManagedInstanceAzureADOnlyAuthenticationsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagedInstanceAzureADOnlyAuthentication); err != nil {
		return ManagedInstanceAzureADOnlyAuthenticationsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByInstancePager - Gets a list of server Azure Active Directory only authentications.
//
// Generated from API version 2023-08-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - managedInstanceName - The name of the managed instance.
//   - options - ManagedInstanceAzureADOnlyAuthenticationsClientListByInstanceOptions contains the optional parameters for the
//     ManagedInstanceAzureADOnlyAuthenticationsClient.NewListByInstancePager method.
func (client *ManagedInstanceAzureADOnlyAuthenticationsClient) NewListByInstancePager(resourceGroupName string, managedInstanceName string, options *ManagedInstanceAzureADOnlyAuthenticationsClientListByInstanceOptions) *runtime.Pager[ManagedInstanceAzureADOnlyAuthenticationsClientListByInstanceResponse] {
	return runtime.NewPager(runtime.PagingHandler[ManagedInstanceAzureADOnlyAuthenticationsClientListByInstanceResponse]{
		More: func(page ManagedInstanceAzureADOnlyAuthenticationsClientListByInstanceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ManagedInstanceAzureADOnlyAuthenticationsClientListByInstanceResponse) (ManagedInstanceAzureADOnlyAuthenticationsClientListByInstanceResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ManagedInstanceAzureADOnlyAuthenticationsClient.NewListByInstancePager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByInstanceCreateRequest(ctx, resourceGroupName, managedInstanceName, options)
			}, nil)
			if err != nil {
				return ManagedInstanceAzureADOnlyAuthenticationsClientListByInstanceResponse{}, err
			}
			return client.listByInstanceHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByInstanceCreateRequest creates the ListByInstance request.
func (client *ManagedInstanceAzureADOnlyAuthenticationsClient) listByInstanceCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, options *ManagedInstanceAzureADOnlyAuthenticationsClientListByInstanceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/azureADOnlyAuthentications"
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
	reqQP.Set("api-version", "2023-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByInstanceHandleResponse handles the ListByInstance response.
func (client *ManagedInstanceAzureADOnlyAuthenticationsClient) listByInstanceHandleResponse(resp *http.Response) (ManagedInstanceAzureADOnlyAuthenticationsClientListByInstanceResponse, error) {
	result := ManagedInstanceAzureADOnlyAuthenticationsClientListByInstanceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagedInstanceAzureADOnlyAuthListResult); err != nil {
		return ManagedInstanceAzureADOnlyAuthenticationsClientListByInstanceResponse{}, err
	}
	return result, nil
}
