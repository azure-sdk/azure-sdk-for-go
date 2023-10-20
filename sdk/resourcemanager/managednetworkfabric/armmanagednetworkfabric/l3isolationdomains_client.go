//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmanagednetworkfabric

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

// L3IsolationDomainsClient contains the methods for the L3IsolationDomains group.
// Don't use this type directly, use NewL3IsolationDomainsClient() instead.
type L3IsolationDomainsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewL3IsolationDomainsClient creates a new instance of L3IsolationDomainsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewL3IsolationDomainsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*L3IsolationDomainsClient, error) {
	cl, err := arm.NewClient(moduleName+".L3IsolationDomainsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &L3IsolationDomainsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCommitConfiguration - Commits the configuration of the given resources.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - l3IsolationDomainName - Name of the L3 Isolation Domain.
//   - options - L3IsolationDomainsClientBeginCommitConfigurationOptions contains the optional parameters for the L3IsolationDomainsClient.BeginCommitConfiguration
//     method.
func (client *L3IsolationDomainsClient) BeginCommitConfiguration(ctx context.Context, resourceGroupName string, l3IsolationDomainName string, options *L3IsolationDomainsClientBeginCommitConfigurationOptions) (*runtime.Poller[L3IsolationDomainsClientCommitConfigurationResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.commitConfiguration(ctx, resourceGroupName, l3IsolationDomainName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[L3IsolationDomainsClientCommitConfigurationResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[L3IsolationDomainsClientCommitConfigurationResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CommitConfiguration - Commits the configuration of the given resources.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
func (client *L3IsolationDomainsClient) commitConfiguration(ctx context.Context, resourceGroupName string, l3IsolationDomainName string, options *L3IsolationDomainsClientBeginCommitConfigurationOptions) (*http.Response, error) {
	var err error
	req, err := client.commitConfigurationCreateRequest(ctx, resourceGroupName, l3IsolationDomainName, options)
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

// commitConfigurationCreateRequest creates the CommitConfiguration request.
func (client *L3IsolationDomainsClient) commitConfigurationCreateRequest(ctx context.Context, resourceGroupName string, l3IsolationDomainName string, options *L3IsolationDomainsClientBeginCommitConfigurationOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetworkFabric/l3IsolationDomains/{l3IsolationDomainName}/commitConfiguration"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if l3IsolationDomainName == "" {
		return nil, errors.New("parameter l3IsolationDomainName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{l3IsolationDomainName}", url.PathEscape(l3IsolationDomainName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginCreate - Create isolation domain resources for layer 3 connectivity between compute nodes and for communication with
// external services .This configuration is applied on the devices only after the creation of
// networks is completed and isolation domain is enabled.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - l3IsolationDomainName - Name of the L3 Isolation Domain.
//   - body - Request payload.
//   - options - L3IsolationDomainsClientBeginCreateOptions contains the optional parameters for the L3IsolationDomainsClient.BeginCreate
//     method.
func (client *L3IsolationDomainsClient) BeginCreate(ctx context.Context, resourceGroupName string, l3IsolationDomainName string, body L3IsolationDomain, options *L3IsolationDomainsClientBeginCreateOptions) (*runtime.Poller[L3IsolationDomainsClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, l3IsolationDomainName, body, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[L3IsolationDomainsClientCreateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[L3IsolationDomainsClientCreateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Create - Create isolation domain resources for layer 3 connectivity between compute nodes and for communication with external
// services .This configuration is applied on the devices only after the creation of
// networks is completed and isolation domain is enabled.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
func (client *L3IsolationDomainsClient) create(ctx context.Context, resourceGroupName string, l3IsolationDomainName string, body L3IsolationDomain, options *L3IsolationDomainsClientBeginCreateOptions) (*http.Response, error) {
	var err error
	req, err := client.createCreateRequest(ctx, resourceGroupName, l3IsolationDomainName, body, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createCreateRequest creates the Create request.
func (client *L3IsolationDomainsClient) createCreateRequest(ctx context.Context, resourceGroupName string, l3IsolationDomainName string, body L3IsolationDomain, options *L3IsolationDomainsClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetworkFabric/l3IsolationDomains/{l3IsolationDomainName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if l3IsolationDomainName == "" {
		return nil, errors.New("parameter l3IsolationDomainName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{l3IsolationDomainName}", url.PathEscape(l3IsolationDomainName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes layer 3 connectivity between compute nodes by managed by named L3 Isolation name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - l3IsolationDomainName - Name of the L3 Isolation Domain.
//   - options - L3IsolationDomainsClientBeginDeleteOptions contains the optional parameters for the L3IsolationDomainsClient.BeginDelete
//     method.
func (client *L3IsolationDomainsClient) BeginDelete(ctx context.Context, resourceGroupName string, l3IsolationDomainName string, options *L3IsolationDomainsClientBeginDeleteOptions) (*runtime.Poller[L3IsolationDomainsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, l3IsolationDomainName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[L3IsolationDomainsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[L3IsolationDomainsClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Deletes layer 3 connectivity between compute nodes by managed by named L3 Isolation name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
func (client *L3IsolationDomainsClient) deleteOperation(ctx context.Context, resourceGroupName string, l3IsolationDomainName string, options *L3IsolationDomainsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, l3IsolationDomainName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *L3IsolationDomainsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, l3IsolationDomainName string, options *L3IsolationDomainsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetworkFabric/l3IsolationDomains/{l3IsolationDomainName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if l3IsolationDomainName == "" {
		return nil, errors.New("parameter l3IsolationDomainName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{l3IsolationDomainName}", url.PathEscape(l3IsolationDomainName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Retrieves details of this L3 Isolation Domain.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - l3IsolationDomainName - Name of the L3 Isolation Domain.
//   - options - L3IsolationDomainsClientGetOptions contains the optional parameters for the L3IsolationDomainsClient.Get method.
func (client *L3IsolationDomainsClient) Get(ctx context.Context, resourceGroupName string, l3IsolationDomainName string, options *L3IsolationDomainsClientGetOptions) (L3IsolationDomainsClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, l3IsolationDomainName, options)
	if err != nil {
		return L3IsolationDomainsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return L3IsolationDomainsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return L3IsolationDomainsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *L3IsolationDomainsClient) getCreateRequest(ctx context.Context, resourceGroupName string, l3IsolationDomainName string, options *L3IsolationDomainsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetworkFabric/l3IsolationDomains/{l3IsolationDomainName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if l3IsolationDomainName == "" {
		return nil, errors.New("parameter l3IsolationDomainName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{l3IsolationDomainName}", url.PathEscape(l3IsolationDomainName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *L3IsolationDomainsClient) getHandleResponse(resp *http.Response) (L3IsolationDomainsClientGetResponse, error) {
	result := L3IsolationDomainsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.L3IsolationDomain); err != nil {
		return L3IsolationDomainsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Displays L3IsolationDomains list by resource group GET method.
//
// Generated from API version 2023-06-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - L3IsolationDomainsClientListByResourceGroupOptions contains the optional parameters for the L3IsolationDomainsClient.NewListByResourceGroupPager
//     method.
func (client *L3IsolationDomainsClient) NewListByResourceGroupPager(resourceGroupName string, options *L3IsolationDomainsClientListByResourceGroupOptions) *runtime.Pager[L3IsolationDomainsClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[L3IsolationDomainsClientListByResourceGroupResponse]{
		More: func(page L3IsolationDomainsClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *L3IsolationDomainsClientListByResourceGroupResponse) (L3IsolationDomainsClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return L3IsolationDomainsClientListByResourceGroupResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return L3IsolationDomainsClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return L3IsolationDomainsClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *L3IsolationDomainsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *L3IsolationDomainsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetworkFabric/l3IsolationDomains"
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
	reqQP.Set("api-version", "2023-06-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *L3IsolationDomainsClient) listByResourceGroupHandleResponse(resp *http.Response) (L3IsolationDomainsClientListByResourceGroupResponse, error) {
	result := L3IsolationDomainsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.L3IsolationDomainsListResult); err != nil {
		return L3IsolationDomainsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - Displays L3IsolationDomains list by subscription GET method.
//
// Generated from API version 2023-06-15
//   - options - L3IsolationDomainsClientListBySubscriptionOptions contains the optional parameters for the L3IsolationDomainsClient.NewListBySubscriptionPager
//     method.
func (client *L3IsolationDomainsClient) NewListBySubscriptionPager(options *L3IsolationDomainsClientListBySubscriptionOptions) *runtime.Pager[L3IsolationDomainsClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[L3IsolationDomainsClientListBySubscriptionResponse]{
		More: func(page L3IsolationDomainsClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *L3IsolationDomainsClientListBySubscriptionResponse) (L3IsolationDomainsClientListBySubscriptionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySubscriptionCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return L3IsolationDomainsClientListBySubscriptionResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return L3IsolationDomainsClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return L3IsolationDomainsClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *L3IsolationDomainsClient) listBySubscriptionCreateRequest(ctx context.Context, options *L3IsolationDomainsClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ManagedNetworkFabric/l3IsolationDomains"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *L3IsolationDomainsClient) listBySubscriptionHandleResponse(resp *http.Response) (L3IsolationDomainsClientListBySubscriptionResponse, error) {
	result := L3IsolationDomainsClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.L3IsolationDomainsListResult); err != nil {
		return L3IsolationDomainsClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// BeginUpdate - API to update certain properties of the L3 Isolation Domain resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - l3IsolationDomainName - Name of the L3 Isolation Domain.
//   - body - API to update certain properties of the L3 Isolation Domain resource.
//   - options - L3IsolationDomainsClientBeginUpdateOptions contains the optional parameters for the L3IsolationDomainsClient.BeginUpdate
//     method.
func (client *L3IsolationDomainsClient) BeginUpdate(ctx context.Context, resourceGroupName string, l3IsolationDomainName string, body L3IsolationDomainPatch, options *L3IsolationDomainsClientBeginUpdateOptions) (*runtime.Poller[L3IsolationDomainsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, l3IsolationDomainName, body, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[L3IsolationDomainsClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[L3IsolationDomainsClientUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Update - API to update certain properties of the L3 Isolation Domain resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
func (client *L3IsolationDomainsClient) update(ctx context.Context, resourceGroupName string, l3IsolationDomainName string, body L3IsolationDomainPatch, options *L3IsolationDomainsClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	req, err := client.updateCreateRequest(ctx, resourceGroupName, l3IsolationDomainName, body, options)
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
func (client *L3IsolationDomainsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, l3IsolationDomainName string, body L3IsolationDomainPatch, options *L3IsolationDomainsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetworkFabric/l3IsolationDomains/{l3IsolationDomainName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if l3IsolationDomainName == "" {
		return nil, errors.New("parameter l3IsolationDomainName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{l3IsolationDomainName}", url.PathEscape(l3IsolationDomainName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginUpdateAdministrativeState - Enables racks for this Isolation Domain.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - l3IsolationDomainName - Name of the L3 Isolation Domain.
//   - body - Request payload.
//   - options - L3IsolationDomainsClientBeginUpdateAdministrativeStateOptions contains the optional parameters for the L3IsolationDomainsClient.BeginUpdateAdministrativeState
//     method.
func (client *L3IsolationDomainsClient) BeginUpdateAdministrativeState(ctx context.Context, resourceGroupName string, l3IsolationDomainName string, body UpdateAdministrativeState, options *L3IsolationDomainsClientBeginUpdateAdministrativeStateOptions) (*runtime.Poller[L3IsolationDomainsClientUpdateAdministrativeStateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.updateAdministrativeState(ctx, resourceGroupName, l3IsolationDomainName, body, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[L3IsolationDomainsClientUpdateAdministrativeStateResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[L3IsolationDomainsClientUpdateAdministrativeStateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// UpdateAdministrativeState - Enables racks for this Isolation Domain.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
func (client *L3IsolationDomainsClient) updateAdministrativeState(ctx context.Context, resourceGroupName string, l3IsolationDomainName string, body UpdateAdministrativeState, options *L3IsolationDomainsClientBeginUpdateAdministrativeStateOptions) (*http.Response, error) {
	var err error
	req, err := client.updateAdministrativeStateCreateRequest(ctx, resourceGroupName, l3IsolationDomainName, body, options)
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

// updateAdministrativeStateCreateRequest creates the UpdateAdministrativeState request.
func (client *L3IsolationDomainsClient) updateAdministrativeStateCreateRequest(ctx context.Context, resourceGroupName string, l3IsolationDomainName string, body UpdateAdministrativeState, options *L3IsolationDomainsClientBeginUpdateAdministrativeStateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetworkFabric/l3IsolationDomains/{l3IsolationDomainName}/updateAdministrativeState"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if l3IsolationDomainName == "" {
		return nil, errors.New("parameter l3IsolationDomainName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{l3IsolationDomainName}", url.PathEscape(l3IsolationDomainName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginValidateConfiguration - Validates the configuration of the resources.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - l3IsolationDomainName - Name of the L3 Isolation Domain.
//   - options - L3IsolationDomainsClientBeginValidateConfigurationOptions contains the optional parameters for the L3IsolationDomainsClient.BeginValidateConfiguration
//     method.
func (client *L3IsolationDomainsClient) BeginValidateConfiguration(ctx context.Context, resourceGroupName string, l3IsolationDomainName string, options *L3IsolationDomainsClientBeginValidateConfigurationOptions) (*runtime.Poller[L3IsolationDomainsClientValidateConfigurationResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.validateConfiguration(ctx, resourceGroupName, l3IsolationDomainName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[L3IsolationDomainsClientValidateConfigurationResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[L3IsolationDomainsClientValidateConfigurationResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// ValidateConfiguration - Validates the configuration of the resources.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
func (client *L3IsolationDomainsClient) validateConfiguration(ctx context.Context, resourceGroupName string, l3IsolationDomainName string, options *L3IsolationDomainsClientBeginValidateConfigurationOptions) (*http.Response, error) {
	var err error
	req, err := client.validateConfigurationCreateRequest(ctx, resourceGroupName, l3IsolationDomainName, options)
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

// validateConfigurationCreateRequest creates the ValidateConfiguration request.
func (client *L3IsolationDomainsClient) validateConfigurationCreateRequest(ctx context.Context, resourceGroupName string, l3IsolationDomainName string, options *L3IsolationDomainsClientBeginValidateConfigurationOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetworkFabric/l3IsolationDomains/{l3IsolationDomainName}/validateConfiguration"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if l3IsolationDomainName == "" {
		return nil, errors.New("parameter l3IsolationDomainName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{l3IsolationDomainName}", url.PathEscape(l3IsolationDomainName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}
