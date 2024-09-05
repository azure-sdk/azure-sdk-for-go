//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

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

// DdosProtectionPlansClient contains the methods for the DdosProtectionPlans group.
// Don't use this type directly, use NewDdosProtectionPlansClient() instead.
type DdosProtectionPlansClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewDdosProtectionPlansClient creates a new instance of DdosProtectionPlansClient with the specified values.
//   - subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
//     ID forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewDdosProtectionPlansClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DdosProtectionPlansClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &DdosProtectionPlansClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates a DDoS protection plan.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-01
//   - resourceGroupName - The name of the resource group.
//   - ddosProtectionPlanName - The name of the DDoS protection plan.
//   - parameters - Parameters supplied to the create or update operation.
//   - options - DdosProtectionPlansClientBeginCreateOrUpdateOptions contains the optional parameters for the DdosProtectionPlansClient.BeginCreateOrUpdate
//     method.
func (client *DdosProtectionPlansClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, ddosProtectionPlanName string, parameters DdosProtectionPlan, options *DdosProtectionPlansClientBeginCreateOrUpdateOptions) (*runtime.Poller[DdosProtectionPlansClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, ddosProtectionPlanName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[DdosProtectionPlansClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[DdosProtectionPlansClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Creates or updates a DDoS protection plan.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-01
func (client *DdosProtectionPlansClient) createOrUpdate(ctx context.Context, resourceGroupName string, ddosProtectionPlanName string, parameters DdosProtectionPlan, options *DdosProtectionPlansClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "DdosProtectionPlansClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, ddosProtectionPlanName, parameters, options)
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
func (client *DdosProtectionPlansClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, ddosProtectionPlanName string, parameters DdosProtectionPlan, options *DdosProtectionPlansClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ddosProtectionPlans/{ddosProtectionPlanName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if ddosProtectionPlanName == "" {
		return nil, errors.New("parameter ddosProtectionPlanName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ddosProtectionPlanName}", url.PathEscape(ddosProtectionPlanName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes the specified DDoS protection plan.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-01
//   - resourceGroupName - The name of the resource group.
//   - ddosProtectionPlanName - The name of the DDoS protection plan.
//   - options - DdosProtectionPlansClientBeginDeleteOptions contains the optional parameters for the DdosProtectionPlansClient.BeginDelete
//     method.
func (client *DdosProtectionPlansClient) BeginDelete(ctx context.Context, resourceGroupName string, ddosProtectionPlanName string, options *DdosProtectionPlansClientBeginDeleteOptions) (*runtime.Poller[DdosProtectionPlansClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, ddosProtectionPlanName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[DdosProtectionPlansClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[DdosProtectionPlansClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Deletes the specified DDoS protection plan.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-01
func (client *DdosProtectionPlansClient) deleteOperation(ctx context.Context, resourceGroupName string, ddosProtectionPlanName string, options *DdosProtectionPlansClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "DdosProtectionPlansClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, ddosProtectionPlanName, options)
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
func (client *DdosProtectionPlansClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, ddosProtectionPlanName string, options *DdosProtectionPlansClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ddosProtectionPlans/{ddosProtectionPlanName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if ddosProtectionPlanName == "" {
		return nil, errors.New("parameter ddosProtectionPlanName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ddosProtectionPlanName}", url.PathEscape(ddosProtectionPlanName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets information about the specified DDoS protection plan.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-01
//   - resourceGroupName - The name of the resource group.
//   - ddosProtectionPlanName - The name of the DDoS protection plan.
//   - options - DdosProtectionPlansClientGetOptions contains the optional parameters for the DdosProtectionPlansClient.Get method.
func (client *DdosProtectionPlansClient) Get(ctx context.Context, resourceGroupName string, ddosProtectionPlanName string, options *DdosProtectionPlansClientGetOptions) (DdosProtectionPlansClientGetResponse, error) {
	var err error
	const operationName = "DdosProtectionPlansClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, ddosProtectionPlanName, options)
	if err != nil {
		return DdosProtectionPlansClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DdosProtectionPlansClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DdosProtectionPlansClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *DdosProtectionPlansClient) getCreateRequest(ctx context.Context, resourceGroupName string, ddosProtectionPlanName string, options *DdosProtectionPlansClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ddosProtectionPlans/{ddosProtectionPlanName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if ddosProtectionPlanName == "" {
		return nil, errors.New("parameter ddosProtectionPlanName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ddosProtectionPlanName}", url.PathEscape(ddosProtectionPlanName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DdosProtectionPlansClient) getHandleResponse(resp *http.Response) (DdosProtectionPlansClientGetResponse, error) {
	result := DdosProtectionPlansClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DdosProtectionPlan); err != nil {
		return DdosProtectionPlansClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets all DDoS protection plans in a subscription.
//
// Generated from API version 2024-03-01
//   - options - DdosProtectionPlansClientListOptions contains the optional parameters for the DdosProtectionPlansClient.NewListPager
//     method.
func (client *DdosProtectionPlansClient) NewListPager(options *DdosProtectionPlansClientListOptions) *runtime.Pager[DdosProtectionPlansClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[DdosProtectionPlansClientListResponse]{
		More: func(page DdosProtectionPlansClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DdosProtectionPlansClientListResponse) (DdosProtectionPlansClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "DdosProtectionPlansClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return DdosProtectionPlansClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *DdosProtectionPlansClient) listCreateRequest(ctx context.Context, options *DdosProtectionPlansClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/ddosProtectionPlans"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *DdosProtectionPlansClient) listHandleResponse(resp *http.Response) (DdosProtectionPlansClientListResponse, error) {
	result := DdosProtectionPlansClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DdosProtectionPlanListResult); err != nil {
		return DdosProtectionPlansClientListResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Gets all the DDoS protection plans in a resource group.
//
// Generated from API version 2024-03-01
//   - resourceGroupName - The name of the resource group.
//   - options - DdosProtectionPlansClientListByResourceGroupOptions contains the optional parameters for the DdosProtectionPlansClient.NewListByResourceGroupPager
//     method.
func (client *DdosProtectionPlansClient) NewListByResourceGroupPager(resourceGroupName string, options *DdosProtectionPlansClientListByResourceGroupOptions) *runtime.Pager[DdosProtectionPlansClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[DdosProtectionPlansClientListByResourceGroupResponse]{
		More: func(page DdosProtectionPlansClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DdosProtectionPlansClientListByResourceGroupResponse) (DdosProtectionPlansClientListByResourceGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "DdosProtectionPlansClient.NewListByResourceGroupPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			}, nil)
			if err != nil {
				return DdosProtectionPlansClientListByResourceGroupResponse{}, err
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *DdosProtectionPlansClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *DdosProtectionPlansClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ddosProtectionPlans"
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
	reqQP.Set("api-version", "2024-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *DdosProtectionPlansClient) listByResourceGroupHandleResponse(resp *http.Response) (DdosProtectionPlansClientListByResourceGroupResponse, error) {
	result := DdosProtectionPlansClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DdosProtectionPlanListResult); err != nil {
		return DdosProtectionPlansClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// UpdateTags - Update a DDoS protection plan tags.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-01
//   - resourceGroupName - The name of the resource group.
//   - ddosProtectionPlanName - The name of the DDoS protection plan.
//   - parameters - Parameters supplied to the update DDoS protection plan resource tags.
//   - options - DdosProtectionPlansClientUpdateTagsOptions contains the optional parameters for the DdosProtectionPlansClient.UpdateTags
//     method.
func (client *DdosProtectionPlansClient) UpdateTags(ctx context.Context, resourceGroupName string, ddosProtectionPlanName string, parameters TagsObject, options *DdosProtectionPlansClientUpdateTagsOptions) (DdosProtectionPlansClientUpdateTagsResponse, error) {
	var err error
	const operationName = "DdosProtectionPlansClient.UpdateTags"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, ddosProtectionPlanName, parameters, options)
	if err != nil {
		return DdosProtectionPlansClientUpdateTagsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DdosProtectionPlansClientUpdateTagsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DdosProtectionPlansClientUpdateTagsResponse{}, err
	}
	resp, err := client.updateTagsHandleResponse(httpResp)
	return resp, err
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *DdosProtectionPlansClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, ddosProtectionPlanName string, parameters TagsObject, options *DdosProtectionPlansClientUpdateTagsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ddosProtectionPlans/{ddosProtectionPlanName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if ddosProtectionPlanName == "" {
		return nil, errors.New("parameter ddosProtectionPlanName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ddosProtectionPlanName}", url.PathEscape(ddosProtectionPlanName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// updateTagsHandleResponse handles the UpdateTags response.
func (client *DdosProtectionPlansClient) updateTagsHandleResponse(resp *http.Response) (DdosProtectionPlansClientUpdateTagsResponse, error) {
	result := DdosProtectionPlansClientUpdateTagsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DdosProtectionPlan); err != nil {
		return DdosProtectionPlansClientUpdateTagsResponse{}, err
	}
	return result, nil
}
