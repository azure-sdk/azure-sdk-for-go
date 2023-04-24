//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armhybridcompute

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

// PrivateLinkScopesClient contains the methods for the PrivateLinkScopes group.
// Don't use this type directly, use NewPrivateLinkScopesClient() instead.
type PrivateLinkScopesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewPrivateLinkScopesClient creates a new instance of PrivateLinkScopesClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewPrivateLinkScopesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*PrivateLinkScopesClient, error) {
	cl, err := arm.NewClient(moduleName+".PrivateLinkScopesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &PrivateLinkScopesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Creates (or updates) a Azure Arc PrivateLinkScope. Note: You cannot specify a different value for InstrumentationKey
// nor AppId in the Put operation.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-15-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - scopeName - The name of the Azure Arc PrivateLinkScope resource.
//   - parameters - Properties that need to be specified to create or update a Azure Arc for Servers and Clusters PrivateLinkScope.
//   - options - PrivateLinkScopesClientCreateOrUpdateOptions contains the optional parameters for the PrivateLinkScopesClient.CreateOrUpdate
//     method.
func (client *PrivateLinkScopesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, scopeName string, parameters PrivateLinkScope, options *PrivateLinkScopesClientCreateOrUpdateOptions) (PrivateLinkScopesClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, scopeName, parameters, options)
	if err != nil {
		return PrivateLinkScopesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PrivateLinkScopesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return PrivateLinkScopesClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *PrivateLinkScopesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, scopeName string, parameters PrivateLinkScope, options *PrivateLinkScopesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridCompute/privateLinkScopes/{scopeName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if scopeName == "" {
		return nil, errors.New("parameter scopeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scopeName}", url.PathEscape(scopeName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *PrivateLinkScopesClient) createOrUpdateHandleResponse(resp *http.Response) (PrivateLinkScopesClientCreateOrUpdateResponse, error) {
	result := PrivateLinkScopesClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateLinkScope); err != nil {
		return PrivateLinkScopesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// BeginDelete - Deletes a Azure Arc PrivateLinkScope.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-15-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - scopeName - The name of the Azure Arc PrivateLinkScope resource.
//   - options - PrivateLinkScopesClientBeginDeleteOptions contains the optional parameters for the PrivateLinkScopesClient.BeginDelete
//     method.
func (client *PrivateLinkScopesClient) BeginDelete(ctx context.Context, resourceGroupName string, scopeName string, options *PrivateLinkScopesClientBeginDeleteOptions) (*runtime.Poller[PrivateLinkScopesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, scopeName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[PrivateLinkScopesClientDeleteResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[PrivateLinkScopesClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Deletes a Azure Arc PrivateLinkScope.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-15-preview
func (client *PrivateLinkScopesClient) deleteOperation(ctx context.Context, resourceGroupName string, scopeName string, options *PrivateLinkScopesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, scopeName, options)
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
func (client *PrivateLinkScopesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, scopeName string, options *PrivateLinkScopesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridCompute/privateLinkScopes/{scopeName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if scopeName == "" {
		return nil, errors.New("parameter scopeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scopeName}", url.PathEscape(scopeName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Returns a Azure Arc PrivateLinkScope.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-15-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - scopeName - The name of the Azure Arc PrivateLinkScope resource.
//   - options - PrivateLinkScopesClientGetOptions contains the optional parameters for the PrivateLinkScopesClient.Get method.
func (client *PrivateLinkScopesClient) Get(ctx context.Context, resourceGroupName string, scopeName string, options *PrivateLinkScopesClientGetOptions) (PrivateLinkScopesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, scopeName, options)
	if err != nil {
		return PrivateLinkScopesClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PrivateLinkScopesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PrivateLinkScopesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *PrivateLinkScopesClient) getCreateRequest(ctx context.Context, resourceGroupName string, scopeName string, options *PrivateLinkScopesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridCompute/privateLinkScopes/{scopeName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if scopeName == "" {
		return nil, errors.New("parameter scopeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scopeName}", url.PathEscape(scopeName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PrivateLinkScopesClient) getHandleResponse(resp *http.Response) (PrivateLinkScopesClientGetResponse, error) {
	result := PrivateLinkScopesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateLinkScope); err != nil {
		return PrivateLinkScopesClientGetResponse{}, err
	}
	return result, nil
}

// GetValidationDetails - Returns a Azure Arc PrivateLinkScope's validation details.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-15-preview
//   - location - The location of the target resource.
//   - privateLinkScopeID - The id (Guid) of the Azure Arc PrivateLinkScope resource.
//   - options - PrivateLinkScopesClientGetValidationDetailsOptions contains the optional parameters for the PrivateLinkScopesClient.GetValidationDetails
//     method.
func (client *PrivateLinkScopesClient) GetValidationDetails(ctx context.Context, location string, privateLinkScopeID string, options *PrivateLinkScopesClientGetValidationDetailsOptions) (PrivateLinkScopesClientGetValidationDetailsResponse, error) {
	req, err := client.getValidationDetailsCreateRequest(ctx, location, privateLinkScopeID, options)
	if err != nil {
		return PrivateLinkScopesClientGetValidationDetailsResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PrivateLinkScopesClientGetValidationDetailsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PrivateLinkScopesClientGetValidationDetailsResponse{}, runtime.NewResponseError(resp)
	}
	return client.getValidationDetailsHandleResponse(resp)
}

// getValidationDetailsCreateRequest creates the GetValidationDetails request.
func (client *PrivateLinkScopesClient) getValidationDetailsCreateRequest(ctx context.Context, location string, privateLinkScopeID string, options *PrivateLinkScopesClientGetValidationDetailsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.HybridCompute/locations/{location}/privateLinkScopes/{privateLinkScopeId}"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if privateLinkScopeID == "" {
		return nil, errors.New("parameter privateLinkScopeID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateLinkScopeId}", url.PathEscape(privateLinkScopeID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getValidationDetailsHandleResponse handles the GetValidationDetails response.
func (client *PrivateLinkScopesClient) getValidationDetailsHandleResponse(resp *http.Response) (PrivateLinkScopesClientGetValidationDetailsResponse, error) {
	result := PrivateLinkScopesClientGetValidationDetailsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateLinkScopeValidationDetails); err != nil {
		return PrivateLinkScopesClientGetValidationDetailsResponse{}, err
	}
	return result, nil
}

// GetValidationDetailsForMachine - Returns a Azure Arc PrivateLinkScope's validation details for a given machine.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-15-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - machineName - The name of the target machine to get the private link scope validation details for.
//   - options - PrivateLinkScopesClientGetValidationDetailsForMachineOptions contains the optional parameters for the PrivateLinkScopesClient.GetValidationDetailsForMachine
//     method.
func (client *PrivateLinkScopesClient) GetValidationDetailsForMachine(ctx context.Context, resourceGroupName string, machineName string, options *PrivateLinkScopesClientGetValidationDetailsForMachineOptions) (PrivateLinkScopesClientGetValidationDetailsForMachineResponse, error) {
	req, err := client.getValidationDetailsForMachineCreateRequest(ctx, resourceGroupName, machineName, options)
	if err != nil {
		return PrivateLinkScopesClientGetValidationDetailsForMachineResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PrivateLinkScopesClientGetValidationDetailsForMachineResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PrivateLinkScopesClientGetValidationDetailsForMachineResponse{}, runtime.NewResponseError(resp)
	}
	return client.getValidationDetailsForMachineHandleResponse(resp)
}

// getValidationDetailsForMachineCreateRequest creates the GetValidationDetailsForMachine request.
func (client *PrivateLinkScopesClient) getValidationDetailsForMachineCreateRequest(ctx context.Context, resourceGroupName string, machineName string, options *PrivateLinkScopesClientGetValidationDetailsForMachineOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridCompute/machines/{machineName}/privateLinkScopes/current"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if machineName == "" {
		return nil, errors.New("parameter machineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{machineName}", url.PathEscape(machineName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getValidationDetailsForMachineHandleResponse handles the GetValidationDetailsForMachine response.
func (client *PrivateLinkScopesClient) getValidationDetailsForMachineHandleResponse(resp *http.Response) (PrivateLinkScopesClientGetValidationDetailsForMachineResponse, error) {
	result := PrivateLinkScopesClientGetValidationDetailsForMachineResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateLinkScopeValidationDetails); err != nil {
		return PrivateLinkScopesClientGetValidationDetailsForMachineResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets a list of all Azure Arc PrivateLinkScopes within a subscription.
//
// Generated from API version 2023-03-15-preview
//   - options - PrivateLinkScopesClientListOptions contains the optional parameters for the PrivateLinkScopesClient.NewListPager
//     method.
func (client *PrivateLinkScopesClient) NewListPager(options *PrivateLinkScopesClientListOptions) *runtime.Pager[PrivateLinkScopesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[PrivateLinkScopesClientListResponse]{
		More: func(page PrivateLinkScopesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *PrivateLinkScopesClientListResponse) (PrivateLinkScopesClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return PrivateLinkScopesClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return PrivateLinkScopesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return PrivateLinkScopesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *PrivateLinkScopesClient) listCreateRequest(ctx context.Context, options *PrivateLinkScopesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.HybridCompute/privateLinkScopes"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *PrivateLinkScopesClient) listHandleResponse(resp *http.Response) (PrivateLinkScopesClientListResponse, error) {
	result := PrivateLinkScopesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateLinkScopeListResult); err != nil {
		return PrivateLinkScopesClientListResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Gets a list of Azure Arc PrivateLinkScopes within a resource group.
//
// Generated from API version 2023-03-15-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - PrivateLinkScopesClientListByResourceGroupOptions contains the optional parameters for the PrivateLinkScopesClient.NewListByResourceGroupPager
//     method.
func (client *PrivateLinkScopesClient) NewListByResourceGroupPager(resourceGroupName string, options *PrivateLinkScopesClientListByResourceGroupOptions) *runtime.Pager[PrivateLinkScopesClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[PrivateLinkScopesClientListByResourceGroupResponse]{
		More: func(page PrivateLinkScopesClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *PrivateLinkScopesClientListByResourceGroupResponse) (PrivateLinkScopesClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return PrivateLinkScopesClientListByResourceGroupResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return PrivateLinkScopesClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return PrivateLinkScopesClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *PrivateLinkScopesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *PrivateLinkScopesClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridCompute/privateLinkScopes"
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
	reqQP.Set("api-version", "2023-03-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *PrivateLinkScopesClient) listByResourceGroupHandleResponse(resp *http.Response) (PrivateLinkScopesClientListByResourceGroupResponse, error) {
	result := PrivateLinkScopesClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateLinkScopeListResult); err != nil {
		return PrivateLinkScopesClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// UpdateTags - Updates an existing PrivateLinkScope's tags. To update other fields use the CreateOrUpdate method.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-15-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - scopeName - The name of the Azure Arc PrivateLinkScope resource.
//   - privateLinkScopeTags - Updated tag information to set into the PrivateLinkScope instance.
//   - options - PrivateLinkScopesClientUpdateTagsOptions contains the optional parameters for the PrivateLinkScopesClient.UpdateTags
//     method.
func (client *PrivateLinkScopesClient) UpdateTags(ctx context.Context, resourceGroupName string, scopeName string, privateLinkScopeTags TagsResource, options *PrivateLinkScopesClientUpdateTagsOptions) (PrivateLinkScopesClientUpdateTagsResponse, error) {
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, scopeName, privateLinkScopeTags, options)
	if err != nil {
		return PrivateLinkScopesClientUpdateTagsResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PrivateLinkScopesClientUpdateTagsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PrivateLinkScopesClientUpdateTagsResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateTagsHandleResponse(resp)
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *PrivateLinkScopesClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, scopeName string, privateLinkScopeTags TagsResource, options *PrivateLinkScopesClientUpdateTagsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridCompute/privateLinkScopes/{scopeName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if scopeName == "" {
		return nil, errors.New("parameter scopeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scopeName}", url.PathEscape(scopeName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, privateLinkScopeTags)
}

// updateTagsHandleResponse handles the UpdateTags response.
func (client *PrivateLinkScopesClient) updateTagsHandleResponse(resp *http.Response) (PrivateLinkScopesClientUpdateTagsResponse, error) {
	result := PrivateLinkScopesClientUpdateTagsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateLinkScope); err != nil {
		return PrivateLinkScopesClientUpdateTagsResponse{}, err
	}
	return result, nil
}
