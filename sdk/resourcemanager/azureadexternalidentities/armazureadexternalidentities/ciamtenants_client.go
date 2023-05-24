//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armazureadexternalidentities

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

// CIAMTenantsClient contains the methods for the CIAMTenants group.
// Don't use this type directly, use NewCIAMTenantsClient() instead.
type CIAMTenantsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewCIAMTenantsClient creates a new instance of CIAMTenantsClient with the specified values.
//   - subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
//     part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewCIAMTenantsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*CIAMTenantsClient, error) {
	cl, err := arm.NewClient(moduleName+".CIAMTenantsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &CIAMTenantsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreate - Initiates an async request to create both the Azure AD for customers tenant and the corresponding Azure resource
// linked to a subscription.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-17-preview
//   - resourceGroupName - The name of the resource group.
//   - resourceName - The initial domain name of the tenant.
//   - options - CIAMTenantsClientBeginCreateOptions contains the optional parameters for the CIAMTenantsClient.BeginCreate method.
func (client *CIAMTenantsClient) BeginCreate(ctx context.Context, resourceGroupName string, resourceName string, createCIAMTenantRequestBody CreateCIAMTenantRequestBody, options *CIAMTenantsClientBeginCreateOptions) (*runtime.Poller[CIAMTenantsClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, resourceName, createCIAMTenantRequestBody, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[CIAMTenantsClientCreateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[CIAMTenantsClientCreateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Create - Initiates an async request to create both the Azure AD for customers tenant and the corresponding Azure resource
// linked to a subscription.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-17-preview
func (client *CIAMTenantsClient) create(ctx context.Context, resourceGroupName string, resourceName string, createCIAMTenantRequestBody CreateCIAMTenantRequestBody, options *CIAMTenantsClientBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, resourceName, createCIAMTenantRequestBody, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createCreateRequest creates the Create request.
func (client *CIAMTenantsClient) createCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, createCIAMTenantRequestBody CreateCIAMTenantRequestBody, options *CIAMTenantsClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureActiveDirectory/ciamDirectories/{resourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-17-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, createCIAMTenantRequestBody)
}

// BeginDelete - Initiates an async operation to delete the Azure AD for customers tenant and Azure resource. The resource
// deletion can only happen as the last step in the tenant deletion process
// [https://learn.microsoft.com/azure/active-directory/enterprise-users/directory-delete-howto].
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-17-preview
//   - resourceGroupName - The name of the resource group.
//   - resourceName - The initial domain name of the tenant.
//   - options - CIAMTenantsClientBeginDeleteOptions contains the optional parameters for the CIAMTenantsClient.BeginDelete method.
func (client *CIAMTenantsClient) BeginDelete(ctx context.Context, resourceGroupName string, resourceName string, options *CIAMTenantsClientBeginDeleteOptions) (*runtime.Poller[CIAMTenantsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, resourceName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[CIAMTenantsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[CIAMTenantsClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Initiates an async operation to delete the Azure AD for customers tenant and Azure resource. The resource deletion
// can only happen as the last step in the tenant deletion process
// [https://learn.microsoft.com/azure/active-directory/enterprise-users/directory-delete-howto].
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-17-preview
func (client *CIAMTenantsClient) deleteOperation(ctx context.Context, resourceGroupName string, resourceName string, options *CIAMTenantsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, resourceName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *CIAMTenantsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *CIAMTenantsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureActiveDirectory/ciamDirectories/{resourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-17-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get the Azure AD for customers tenant resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-17-preview
//   - resourceGroupName - The name of the resource group.
//   - resourceName - The initial domain name of the tenant.
//   - options - CIAMTenantsClientGetOptions contains the optional parameters for the CIAMTenantsClient.Get method.
func (client *CIAMTenantsClient) Get(ctx context.Context, resourceGroupName string, resourceName string, options *CIAMTenantsClientGetOptions) (CIAMTenantsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, resourceName, options)
	if err != nil {
		return CIAMTenantsClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return CIAMTenantsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return CIAMTenantsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *CIAMTenantsClient) getCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *CIAMTenantsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureActiveDirectory/ciamDirectories/{resourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-17-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *CIAMTenantsClient) getHandleResponse(resp *http.Response) (CIAMTenantsClientGetResponse, error) {
	result := CIAMTenantsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CIAMTenantResource); err != nil {
		return CIAMTenantsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Get all the Azure AD for customers tenants resources in a resource group.
//
// Generated from API version 2023-05-17-preview
//   - resourceGroupName - The name of the resource group.
//   - options - CIAMTenantsClientListByResourceGroupOptions contains the optional parameters for the CIAMTenantsClient.NewListByResourceGroupPager
//     method.
func (client *CIAMTenantsClient) NewListByResourceGroupPager(resourceGroupName string, options *CIAMTenantsClientListByResourceGroupOptions) *runtime.Pager[CIAMTenantsClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[CIAMTenantsClientListByResourceGroupResponse]{
		More: func(page CIAMTenantsClientListByResourceGroupResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *CIAMTenantsClientListByResourceGroupResponse) (CIAMTenantsClientListByResourceGroupResponse, error) {
			req, err := client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			if err != nil {
				return CIAMTenantsClientListByResourceGroupResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return CIAMTenantsClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return CIAMTenantsClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *CIAMTenantsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *CIAMTenantsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureActiveDirectory/ciamDirectories"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-17-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *CIAMTenantsClient) listByResourceGroupHandleResponse(resp *http.Response) (CIAMTenantsClientListByResourceGroupResponse, error) {
	result := CIAMTenantsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CIAMTenantResourceList); err != nil {
		return CIAMTenantsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - Get all the Azure AD for customers tenant resources in a subscription.
//
// Generated from API version 2023-05-17-preview
//   - options - CIAMTenantsClientListBySubscriptionOptions contains the optional parameters for the CIAMTenantsClient.NewListBySubscriptionPager
//     method.
func (client *CIAMTenantsClient) NewListBySubscriptionPager(options *CIAMTenantsClientListBySubscriptionOptions) *runtime.Pager[CIAMTenantsClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[CIAMTenantsClientListBySubscriptionResponse]{
		More: func(page CIAMTenantsClientListBySubscriptionResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *CIAMTenantsClientListBySubscriptionResponse) (CIAMTenantsClientListBySubscriptionResponse, error) {
			req, err := client.listBySubscriptionCreateRequest(ctx, options)
			if err != nil {
				return CIAMTenantsClientListBySubscriptionResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return CIAMTenantsClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return CIAMTenantsClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *CIAMTenantsClient) listBySubscriptionCreateRequest(ctx context.Context, options *CIAMTenantsClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.AzureActiveDirectory/ciamDirectories"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-17-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *CIAMTenantsClient) listBySubscriptionHandleResponse(resp *http.Response) (CIAMTenantsClientListBySubscriptionResponse, error) {
	result := CIAMTenantsClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CIAMTenantResourceList); err != nil {
		return CIAMTenantsClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// Update - Update the Azure AD for customers tenant resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-17-preview
//   - resourceGroupName - The name of the resource group.
//   - resourceName - The initial domain name of the tenant.
//   - options - CIAMTenantsClientUpdateOptions contains the optional parameters for the CIAMTenantsClient.Update method.
func (client *CIAMTenantsClient) Update(ctx context.Context, resourceGroupName string, resourceName string, updateCIAMTenantRequestBody CIAMTenantUpdateRequest, options *CIAMTenantsClientUpdateOptions) (CIAMTenantsClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, resourceName, updateCIAMTenantRequestBody, options)
	if err != nil {
		return CIAMTenantsClientUpdateResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return CIAMTenantsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return CIAMTenantsClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *CIAMTenantsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, updateCIAMTenantRequestBody CIAMTenantUpdateRequest, options *CIAMTenantsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureActiveDirectory/ciamDirectories/{resourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-17-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, updateCIAMTenantRequestBody)
}

// updateHandleResponse handles the Update response.
func (client *CIAMTenantsClient) updateHandleResponse(resp *http.Response) (CIAMTenantsClientUpdateResponse, error) {
	result := CIAMTenantsClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CIAMTenantResource); err != nil {
		return CIAMTenantsClientUpdateResponse{}, err
	}
	return result, nil
}
