//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmobilenetwork

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

// SimPoliciesClient contains the methods for the SimPolicies group.
// Don't use this type directly, use NewSimPoliciesClient() instead.
type SimPoliciesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewSimPoliciesClient creates a new instance of SimPoliciesClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewSimPoliciesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SimPoliciesClient, error) {
	cl, err := arm.NewClient(moduleName+".SimPoliciesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &SimPoliciesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates a SIM policy. Must be created in the same location as its parent mobile network.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - mobileNetworkName - The name of the mobile network.
//   - simPolicyName - The name of the SIM policy.
//   - parameters - Parameters supplied to the create or update SIM policy operation.
//   - options - SimPoliciesClientBeginCreateOrUpdateOptions contains the optional parameters for the SimPoliciesClient.BeginCreateOrUpdate
//     method.
func (client *SimPoliciesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, mobileNetworkName string, simPolicyName string, parameters SimPolicy, options *SimPoliciesClientBeginCreateOrUpdateOptions) (*runtime.Poller[SimPoliciesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, mobileNetworkName, simPolicyName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[SimPoliciesClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[SimPoliciesClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Creates or updates a SIM policy. Must be created in the same location as its parent mobile network.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01
func (client *SimPoliciesClient) createOrUpdate(ctx context.Context, resourceGroupName string, mobileNetworkName string, simPolicyName string, parameters SimPolicy, options *SimPoliciesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, mobileNetworkName, simPolicyName, parameters, options)
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

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *SimPoliciesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, mobileNetworkName string, simPolicyName string, parameters SimPolicy, options *SimPoliciesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MobileNetwork/mobileNetworks/{mobileNetworkName}/simPolicies/{simPolicyName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if mobileNetworkName == "" {
		return nil, errors.New("parameter mobileNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{mobileNetworkName}", url.PathEscape(mobileNetworkName))
	if simPolicyName == "" {
		return nil, errors.New("parameter simPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{simPolicyName}", url.PathEscape(simPolicyName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes the specified SIM policy.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - mobileNetworkName - The name of the mobile network.
//   - simPolicyName - The name of the SIM policy.
//   - options - SimPoliciesClientBeginDeleteOptions contains the optional parameters for the SimPoliciesClient.BeginDelete method.
func (client *SimPoliciesClient) BeginDelete(ctx context.Context, resourceGroupName string, mobileNetworkName string, simPolicyName string, options *SimPoliciesClientBeginDeleteOptions) (*runtime.Poller[SimPoliciesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, mobileNetworkName, simPolicyName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[SimPoliciesClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[SimPoliciesClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Deletes the specified SIM policy.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01
func (client *SimPoliciesClient) deleteOperation(ctx context.Context, resourceGroupName string, mobileNetworkName string, simPolicyName string, options *SimPoliciesClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, mobileNetworkName, simPolicyName, options)
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
func (client *SimPoliciesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, mobileNetworkName string, simPolicyName string, options *SimPoliciesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MobileNetwork/mobileNetworks/{mobileNetworkName}/simPolicies/{simPolicyName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if mobileNetworkName == "" {
		return nil, errors.New("parameter mobileNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{mobileNetworkName}", url.PathEscape(mobileNetworkName))
	if simPolicyName == "" {
		return nil, errors.New("parameter simPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{simPolicyName}", url.PathEscape(simPolicyName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets information about the specified SIM policy.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - mobileNetworkName - The name of the mobile network.
//   - simPolicyName - The name of the SIM policy.
//   - options - SimPoliciesClientGetOptions contains the optional parameters for the SimPoliciesClient.Get method.
func (client *SimPoliciesClient) Get(ctx context.Context, resourceGroupName string, mobileNetworkName string, simPolicyName string, options *SimPoliciesClientGetOptions) (SimPoliciesClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, mobileNetworkName, simPolicyName, options)
	if err != nil {
		return SimPoliciesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SimPoliciesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SimPoliciesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *SimPoliciesClient) getCreateRequest(ctx context.Context, resourceGroupName string, mobileNetworkName string, simPolicyName string, options *SimPoliciesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MobileNetwork/mobileNetworks/{mobileNetworkName}/simPolicies/{simPolicyName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if mobileNetworkName == "" {
		return nil, errors.New("parameter mobileNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{mobileNetworkName}", url.PathEscape(mobileNetworkName))
	if simPolicyName == "" {
		return nil, errors.New("parameter simPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{simPolicyName}", url.PathEscape(simPolicyName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SimPoliciesClient) getHandleResponse(resp *http.Response) (SimPoliciesClientGetResponse, error) {
	result := SimPoliciesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SimPolicy); err != nil {
		return SimPoliciesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByMobileNetworkPager - Gets all the SIM policies in a mobile network.
//
// Generated from API version 2023-06-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - mobileNetworkName - The name of the mobile network.
//   - options - SimPoliciesClientListByMobileNetworkOptions contains the optional parameters for the SimPoliciesClient.NewListByMobileNetworkPager
//     method.
func (client *SimPoliciesClient) NewListByMobileNetworkPager(resourceGroupName string, mobileNetworkName string, options *SimPoliciesClientListByMobileNetworkOptions) *runtime.Pager[SimPoliciesClientListByMobileNetworkResponse] {
	return runtime.NewPager(runtime.PagingHandler[SimPoliciesClientListByMobileNetworkResponse]{
		More: func(page SimPoliciesClientListByMobileNetworkResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SimPoliciesClientListByMobileNetworkResponse) (SimPoliciesClientListByMobileNetworkResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByMobileNetworkCreateRequest(ctx, resourceGroupName, mobileNetworkName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return SimPoliciesClientListByMobileNetworkResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return SimPoliciesClientListByMobileNetworkResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return SimPoliciesClientListByMobileNetworkResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByMobileNetworkHandleResponse(resp)
		},
	})
}

// listByMobileNetworkCreateRequest creates the ListByMobileNetwork request.
func (client *SimPoliciesClient) listByMobileNetworkCreateRequest(ctx context.Context, resourceGroupName string, mobileNetworkName string, options *SimPoliciesClientListByMobileNetworkOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MobileNetwork/mobileNetworks/{mobileNetworkName}/simPolicies"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if mobileNetworkName == "" {
		return nil, errors.New("parameter mobileNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{mobileNetworkName}", url.PathEscape(mobileNetworkName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByMobileNetworkHandleResponse handles the ListByMobileNetwork response.
func (client *SimPoliciesClient) listByMobileNetworkHandleResponse(resp *http.Response) (SimPoliciesClientListByMobileNetworkResponse, error) {
	result := SimPoliciesClientListByMobileNetworkResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SimPolicyListResult); err != nil {
		return SimPoliciesClientListByMobileNetworkResponse{}, err
	}
	return result, nil
}

// UpdateTags - Updates SIM policy tags.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - mobileNetworkName - The name of the mobile network.
//   - simPolicyName - The name of the SIM policy.
//   - parameters - Parameters supplied to update SIM policy tags.
//   - options - SimPoliciesClientUpdateTagsOptions contains the optional parameters for the SimPoliciesClient.UpdateTags method.
func (client *SimPoliciesClient) UpdateTags(ctx context.Context, resourceGroupName string, mobileNetworkName string, simPolicyName string, parameters TagsObject, options *SimPoliciesClientUpdateTagsOptions) (SimPoliciesClientUpdateTagsResponse, error) {
	var err error
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, mobileNetworkName, simPolicyName, parameters, options)
	if err != nil {
		return SimPoliciesClientUpdateTagsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SimPoliciesClientUpdateTagsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SimPoliciesClientUpdateTagsResponse{}, err
	}
	resp, err := client.updateTagsHandleResponse(httpResp)
	return resp, err
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *SimPoliciesClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, mobileNetworkName string, simPolicyName string, parameters TagsObject, options *SimPoliciesClientUpdateTagsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MobileNetwork/mobileNetworks/{mobileNetworkName}/simPolicies/{simPolicyName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if mobileNetworkName == "" {
		return nil, errors.New("parameter mobileNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{mobileNetworkName}", url.PathEscape(mobileNetworkName))
	if simPolicyName == "" {
		return nil, errors.New("parameter simPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{simPolicyName}", url.PathEscape(simPolicyName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// updateTagsHandleResponse handles the UpdateTags response.
func (client *SimPoliciesClient) updateTagsHandleResponse(resp *http.Response) (SimPoliciesClientUpdateTagsResponse, error) {
	result := SimPoliciesClientUpdateTagsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SimPolicy); err != nil {
		return SimPoliciesClientUpdateTagsResponse{}, err
	}
	return result, nil
}
