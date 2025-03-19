// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armfrontdoor

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

// NetworkExperimentProfilesClient contains the methods for the NetworkExperimentProfiles group.
// Don't use this type directly, use NewNetworkExperimentProfilesClient() instead.
type NetworkExperimentProfilesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewNetworkExperimentProfilesClient creates a new instance of NetworkExperimentProfilesClient with the specified values.
//   - subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
//     ID forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewNetworkExperimentProfilesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*NetworkExperimentProfilesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &NetworkExperimentProfilesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates an NetworkExperiment Profile
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-11-01
//   - profileName - The Profile identifier associated with the Tenant and Partner
//   - resourceGroupName - Name of the Resource group within the Azure subscription.
//   - parameters - An Network Experiment Profile
//   - options - NetworkExperimentProfilesClientBeginCreateOrUpdateOptions contains the optional parameters for the NetworkExperimentProfilesClient.BeginCreateOrUpdate
//     method.
func (client *NetworkExperimentProfilesClient) BeginCreateOrUpdate(ctx context.Context, profileName string, resourceGroupName string, parameters Profile, options *NetworkExperimentProfilesClientBeginCreateOrUpdateOptions) (*runtime.Poller[NetworkExperimentProfilesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, profileName, resourceGroupName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[NetworkExperimentProfilesClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[NetworkExperimentProfilesClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Creates an NetworkExperiment Profile
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-11-01
func (client *NetworkExperimentProfilesClient) createOrUpdate(ctx context.Context, profileName string, resourceGroupName string, parameters Profile, options *NetworkExperimentProfilesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "NetworkExperimentProfilesClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, profileName, resourceGroupName, parameters, options)
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
func (client *NetworkExperimentProfilesClient) createOrUpdateCreateRequest(ctx context.Context, profileName string, resourceGroupName string, parameters Profile, _ *NetworkExperimentProfilesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/NetworkExperimentProfiles/{profileName}"
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes an NetworkExperiment Profile by ProfileName
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-11-01
//   - resourceGroupName - Name of the Resource group within the Azure subscription.
//   - profileName - The Profile identifier associated with the Tenant and Partner
//   - options - NetworkExperimentProfilesClientBeginDeleteOptions contains the optional parameters for the NetworkExperimentProfilesClient.BeginDelete
//     method.
func (client *NetworkExperimentProfilesClient) BeginDelete(ctx context.Context, resourceGroupName string, profileName string, options *NetworkExperimentProfilesClientBeginDeleteOptions) (*runtime.Poller[NetworkExperimentProfilesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, profileName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[NetworkExperimentProfilesClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[NetworkExperimentProfilesClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Deletes an NetworkExperiment Profile by ProfileName
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-11-01
func (client *NetworkExperimentProfilesClient) deleteOperation(ctx context.Context, resourceGroupName string, profileName string, options *NetworkExperimentProfilesClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "NetworkExperimentProfilesClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, profileName, options)
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
func (client *NetworkExperimentProfilesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, profileName string, _ *NetworkExperimentProfilesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/NetworkExperimentProfiles/{profileName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets an NetworkExperiment Profile by ProfileName
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-11-01
//   - resourceGroupName - Name of the Resource group within the Azure subscription.
//   - profileName - The Profile identifier associated with the Tenant and Partner
//   - options - NetworkExperimentProfilesClientGetOptions contains the optional parameters for the NetworkExperimentProfilesClient.Get
//     method.
func (client *NetworkExperimentProfilesClient) Get(ctx context.Context, resourceGroupName string, profileName string, options *NetworkExperimentProfilesClientGetOptions) (NetworkExperimentProfilesClientGetResponse, error) {
	var err error
	const operationName = "NetworkExperimentProfilesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, profileName, options)
	if err != nil {
		return NetworkExperimentProfilesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return NetworkExperimentProfilesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return NetworkExperimentProfilesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *NetworkExperimentProfilesClient) getCreateRequest(ctx context.Context, resourceGroupName string, profileName string, _ *NetworkExperimentProfilesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/NetworkExperimentProfiles/{profileName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *NetworkExperimentProfilesClient) getHandleResponse(resp *http.Response) (NetworkExperimentProfilesClientGetResponse, error) {
	result := NetworkExperimentProfilesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Profile); err != nil {
		return NetworkExperimentProfilesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets a list of Network Experiment Profiles under a subscription
//
// Generated from API version 2019-11-01
//   - options - NetworkExperimentProfilesClientListOptions contains the optional parameters for the NetworkExperimentProfilesClient.NewListPager
//     method.
func (client *NetworkExperimentProfilesClient) NewListPager(options *NetworkExperimentProfilesClientListOptions) *runtime.Pager[NetworkExperimentProfilesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[NetworkExperimentProfilesClientListResponse]{
		More: func(page NetworkExperimentProfilesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *NetworkExperimentProfilesClientListResponse) (NetworkExperimentProfilesClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "NetworkExperimentProfilesClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return NetworkExperimentProfilesClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *NetworkExperimentProfilesClient) listCreateRequest(ctx context.Context, _ *NetworkExperimentProfilesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/NetworkExperimentProfiles"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *NetworkExperimentProfilesClient) listHandleResponse(resp *http.Response) (NetworkExperimentProfilesClientListResponse, error) {
	result := NetworkExperimentProfilesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProfileList); err != nil {
		return NetworkExperimentProfilesClientListResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Gets a list of Network Experiment Profiles within a resource group under a subscription
//
// Generated from API version 2019-11-01
//   - resourceGroupName - Name of the Resource group within the Azure subscription.
//   - options - NetworkExperimentProfilesClientListByResourceGroupOptions contains the optional parameters for the NetworkExperimentProfilesClient.NewListByResourceGroupPager
//     method.
func (client *NetworkExperimentProfilesClient) NewListByResourceGroupPager(resourceGroupName string, options *NetworkExperimentProfilesClientListByResourceGroupOptions) *runtime.Pager[NetworkExperimentProfilesClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[NetworkExperimentProfilesClientListByResourceGroupResponse]{
		More: func(page NetworkExperimentProfilesClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *NetworkExperimentProfilesClientListByResourceGroupResponse) (NetworkExperimentProfilesClientListByResourceGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "NetworkExperimentProfilesClient.NewListByResourceGroupPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			}, nil)
			if err != nil {
				return NetworkExperimentProfilesClientListByResourceGroupResponse{}, err
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *NetworkExperimentProfilesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, _ *NetworkExperimentProfilesClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/NetworkExperimentProfiles"
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
	reqQP.Set("api-version", "2019-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *NetworkExperimentProfilesClient) listByResourceGroupHandleResponse(resp *http.Response) (NetworkExperimentProfilesClientListByResourceGroupResponse, error) {
	result := NetworkExperimentProfilesClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProfileList); err != nil {
		return NetworkExperimentProfilesClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Updates an NetworkExperimentProfiles
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-11-01
//   - resourceGroupName - Name of the Resource group within the Azure subscription.
//   - profileName - The Profile identifier associated with the Tenant and Partner
//   - parameters - The Profile Update Model
//   - options - NetworkExperimentProfilesClientBeginUpdateOptions contains the optional parameters for the NetworkExperimentProfilesClient.BeginUpdate
//     method.
func (client *NetworkExperimentProfilesClient) BeginUpdate(ctx context.Context, resourceGroupName string, profileName string, parameters ProfileUpdateModel, options *NetworkExperimentProfilesClientBeginUpdateOptions) (*runtime.Poller[NetworkExperimentProfilesClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, profileName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[NetworkExperimentProfilesClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[NetworkExperimentProfilesClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Update - Updates an NetworkExperimentProfiles
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-11-01
func (client *NetworkExperimentProfilesClient) update(ctx context.Context, resourceGroupName string, profileName string, parameters ProfileUpdateModel, options *NetworkExperimentProfilesClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "NetworkExperimentProfilesClient.BeginUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, profileName, parameters, options)
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
func (client *NetworkExperimentProfilesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, profileName string, parameters ProfileUpdateModel, _ *NetworkExperimentProfilesClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/NetworkExperimentProfiles/{profileName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}
