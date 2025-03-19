// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armavs

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

// PlacementPoliciesClient contains the methods for the PlacementPolicies group.
// Don't use this type directly, use NewPlacementPoliciesClient() instead.
type PlacementPoliciesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewPlacementPoliciesClient creates a new instance of PlacementPoliciesClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewPlacementPoliciesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*PlacementPoliciesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &PlacementPoliciesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create a PlacementPolicy
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - privateCloudName - Name of the private cloud
//   - clusterName - Name of the cluster
//   - placementPolicyName - Name of the placement policy.
//   - placementPolicy - Resource create parameters.
//   - options - PlacementPoliciesClientBeginCreateOrUpdateOptions contains the optional parameters for the PlacementPoliciesClient.BeginCreateOrUpdate
//     method.
func (client *PlacementPoliciesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, privateCloudName string, clusterName string, placementPolicyName string, placementPolicy PlacementPolicy, options *PlacementPoliciesClientBeginCreateOrUpdateOptions) (*runtime.Poller[PlacementPoliciesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, privateCloudName, clusterName, placementPolicyName, placementPolicy, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[PlacementPoliciesClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[PlacementPoliciesClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Create a PlacementPolicy
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-09-01
func (client *PlacementPoliciesClient) createOrUpdate(ctx context.Context, resourceGroupName string, privateCloudName string, clusterName string, placementPolicyName string, placementPolicy PlacementPolicy, options *PlacementPoliciesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "PlacementPoliciesClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, privateCloudName, clusterName, placementPolicyName, placementPolicy, options)
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
func (client *PlacementPoliciesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, privateCloudName string, clusterName string, placementPolicyName string, placementPolicy PlacementPolicy, _ *PlacementPoliciesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AVS/privateClouds/{privateCloudName}/clusters/{clusterName}/placementPolicies/{placementPolicyName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if privateCloudName == "" {
		return nil, errors.New("parameter privateCloudName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateCloudName}", url.PathEscape(privateCloudName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if placementPolicyName == "" {
		return nil, errors.New("parameter placementPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{placementPolicyName}", url.PathEscape(placementPolicyName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, placementPolicy); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Delete a PlacementPolicy
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - privateCloudName - Name of the private cloud
//   - clusterName - Name of the cluster
//   - placementPolicyName - Name of the placement policy.
//   - options - PlacementPoliciesClientBeginDeleteOptions contains the optional parameters for the PlacementPoliciesClient.BeginDelete
//     method.
func (client *PlacementPoliciesClient) BeginDelete(ctx context.Context, resourceGroupName string, privateCloudName string, clusterName string, placementPolicyName string, options *PlacementPoliciesClientBeginDeleteOptions) (*runtime.Poller[PlacementPoliciesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, privateCloudName, clusterName, placementPolicyName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[PlacementPoliciesClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[PlacementPoliciesClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Delete a PlacementPolicy
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-09-01
func (client *PlacementPoliciesClient) deleteOperation(ctx context.Context, resourceGroupName string, privateCloudName string, clusterName string, placementPolicyName string, options *PlacementPoliciesClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "PlacementPoliciesClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, privateCloudName, clusterName, placementPolicyName, options)
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
func (client *PlacementPoliciesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, privateCloudName string, clusterName string, placementPolicyName string, _ *PlacementPoliciesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AVS/privateClouds/{privateCloudName}/clusters/{clusterName}/placementPolicies/{placementPolicyName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if privateCloudName == "" {
		return nil, errors.New("parameter privateCloudName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateCloudName}", url.PathEscape(privateCloudName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if placementPolicyName == "" {
		return nil, errors.New("parameter placementPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{placementPolicyName}", url.PathEscape(placementPolicyName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get a PlacementPolicy
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - privateCloudName - Name of the private cloud
//   - clusterName - Name of the cluster
//   - placementPolicyName - Name of the placement policy.
//   - options - PlacementPoliciesClientGetOptions contains the optional parameters for the PlacementPoliciesClient.Get method.
func (client *PlacementPoliciesClient) Get(ctx context.Context, resourceGroupName string, privateCloudName string, clusterName string, placementPolicyName string, options *PlacementPoliciesClientGetOptions) (PlacementPoliciesClientGetResponse, error) {
	var err error
	const operationName = "PlacementPoliciesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, privateCloudName, clusterName, placementPolicyName, options)
	if err != nil {
		return PlacementPoliciesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PlacementPoliciesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PlacementPoliciesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *PlacementPoliciesClient) getCreateRequest(ctx context.Context, resourceGroupName string, privateCloudName string, clusterName string, placementPolicyName string, _ *PlacementPoliciesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AVS/privateClouds/{privateCloudName}/clusters/{clusterName}/placementPolicies/{placementPolicyName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if privateCloudName == "" {
		return nil, errors.New("parameter privateCloudName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateCloudName}", url.PathEscape(privateCloudName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if placementPolicyName == "" {
		return nil, errors.New("parameter placementPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{placementPolicyName}", url.PathEscape(placementPolicyName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PlacementPoliciesClient) getHandleResponse(resp *http.Response) (PlacementPoliciesClientGetResponse, error) {
	result := PlacementPoliciesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PlacementPolicy); err != nil {
		return PlacementPoliciesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List PlacementPolicy resources by Cluster
//
// Generated from API version 2024-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - privateCloudName - Name of the private cloud
//   - clusterName - Name of the cluster
//   - options - PlacementPoliciesClientListOptions contains the optional parameters for the PlacementPoliciesClient.NewListPager
//     method.
func (client *PlacementPoliciesClient) NewListPager(resourceGroupName string, privateCloudName string, clusterName string, options *PlacementPoliciesClientListOptions) *runtime.Pager[PlacementPoliciesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[PlacementPoliciesClientListResponse]{
		More: func(page PlacementPoliciesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *PlacementPoliciesClientListResponse) (PlacementPoliciesClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "PlacementPoliciesClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, privateCloudName, clusterName, options)
			}, nil)
			if err != nil {
				return PlacementPoliciesClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *PlacementPoliciesClient) listCreateRequest(ctx context.Context, resourceGroupName string, privateCloudName string, clusterName string, _ *PlacementPoliciesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AVS/privateClouds/{privateCloudName}/clusters/{clusterName}/placementPolicies"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if privateCloudName == "" {
		return nil, errors.New("parameter privateCloudName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateCloudName}", url.PathEscape(privateCloudName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *PlacementPoliciesClient) listHandleResponse(resp *http.Response) (PlacementPoliciesClientListResponse, error) {
	result := PlacementPoliciesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PlacementPoliciesList); err != nil {
		return PlacementPoliciesClientListResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Update a PlacementPolicy
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - privateCloudName - Name of the private cloud
//   - clusterName - Name of the cluster
//   - placementPolicyName - Name of the placement policy.
//   - placementPolicyUpdate - The placement policy properties to be updated.
//   - options - PlacementPoliciesClientBeginUpdateOptions contains the optional parameters for the PlacementPoliciesClient.BeginUpdate
//     method.
func (client *PlacementPoliciesClient) BeginUpdate(ctx context.Context, resourceGroupName string, privateCloudName string, clusterName string, placementPolicyName string, placementPolicyUpdate PlacementPolicyUpdate, options *PlacementPoliciesClientBeginUpdateOptions) (*runtime.Poller[PlacementPoliciesClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, privateCloudName, clusterName, placementPolicyName, placementPolicyUpdate, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[PlacementPoliciesClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[PlacementPoliciesClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Update - Update a PlacementPolicy
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-09-01
func (client *PlacementPoliciesClient) update(ctx context.Context, resourceGroupName string, privateCloudName string, clusterName string, placementPolicyName string, placementPolicyUpdate PlacementPolicyUpdate, options *PlacementPoliciesClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "PlacementPoliciesClient.BeginUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, privateCloudName, clusterName, placementPolicyName, placementPolicyUpdate, options)
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
func (client *PlacementPoliciesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, privateCloudName string, clusterName string, placementPolicyName string, placementPolicyUpdate PlacementPolicyUpdate, _ *PlacementPoliciesClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AVS/privateClouds/{privateCloudName}/clusters/{clusterName}/placementPolicies/{placementPolicyName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if privateCloudName == "" {
		return nil, errors.New("parameter privateCloudName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateCloudName}", url.PathEscape(privateCloudName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if placementPolicyName == "" {
		return nil, errors.New("parameter placementPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{placementPolicyName}", url.PathEscape(placementPolicyName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, placementPolicyUpdate); err != nil {
		return nil, err
	}
	return req, nil
}
