//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armquota

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

// GroupQuotasClient contains the methods for the GroupQuotas group.
// Don't use this type directly, use NewGroupQuotasClient() instead.
type GroupQuotasClient struct {
	internal *arm.Client
}

// NewGroupQuotasClient creates a new instance of GroupQuotasClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewGroupQuotasClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*GroupQuotasClient, error) {
	cl, err := arm.NewClient(moduleName+".GroupQuotasClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &GroupQuotasClient{
		internal: cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates a new GroupQuota for the name passed. A RequestId will be returned by the Service. The status
// can be polled periodically. The status Async polling is using standards defined at -
// https://github.com/Azure/azure-resource-manager-rpc/blob/master/v1.0/async-api-reference.md#asynchronous-operations. Use
// the OperationsStatus URI provided in Azure-AsyncOperation header, the duration
// will be specified in retry-after header. Once the operation gets to terminal state - Succeeded | Failed, then the URI will
// change to Get URI and full details can be checked.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01-preview
//   - mgID - Management Group Id.
//   - groupQuotaName - The GroupQuota name. The name should be unique for the provided context tenantId/MgId.
//   - options - GroupQuotasClientBeginCreateOrUpdateOptions contains the optional parameters for the GroupQuotasClient.BeginCreateOrUpdate
//     method.
func (client *GroupQuotasClient) BeginCreateOrUpdate(ctx context.Context, mgID string, groupQuotaName string, options *GroupQuotasClientBeginCreateOrUpdateOptions) (*runtime.Poller[GroupQuotasClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, mgID, groupQuotaName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[GroupQuotasClientCreateOrUpdateResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[GroupQuotasClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Creates a new GroupQuota for the name passed. A RequestId will be returned by the Service. The status
// can be polled periodically. The status Async polling is using standards defined at -
// https://github.com/Azure/azure-resource-manager-rpc/blob/master/v1.0/async-api-reference.md#asynchronous-operations. Use
// the OperationsStatus URI provided in Azure-AsyncOperation header, the duration
// will be specified in retry-after header. Once the operation gets to terminal state - Succeeded | Failed, then the URI will
// change to Get URI and full details can be checked.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01-preview
func (client *GroupQuotasClient) createOrUpdate(ctx context.Context, mgID string, groupQuotaName string, options *GroupQuotasClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	req, err := client.createOrUpdateCreateRequest(ctx, mgID, groupQuotaName, options)
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
func (client *GroupQuotasClient) createOrUpdateCreateRequest(ctx context.Context, mgID string, groupQuotaName string, options *GroupQuotasClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Management/managementGroups/{mgId}/providers/Microsoft.Quota/groupQuotas/{groupQuotaName}"
	if mgID == "" {
		return nil, errors.New("parameter mgID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{mgId}", url.PathEscape(mgID))
	if groupQuotaName == "" {
		return nil, errors.New("parameter groupQuotaName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupQuotaName}", url.PathEscape(groupQuotaName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.GroupQuotaPutRequestBody != nil {
		if err := runtime.MarshalAsJSON(req, *options.GroupQuotaPutRequestBody); err != nil {
			return nil, err
		}
		return req, nil
	}
	return req, nil
}

// Delete - Deletes the GroupQuotas for the name passed. All the remaining shareQuota in the GroupQuotas will be lost.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01-preview
//   - mgID - Management Group Id.
//   - groupQuotaName - The GroupQuota name. The name should be unique for the provided context tenantId/MgId.
//   - options - GroupQuotasClientDeleteOptions contains the optional parameters for the GroupQuotasClient.Delete method.
func (client *GroupQuotasClient) Delete(ctx context.Context, mgID string, groupQuotaName string, options *GroupQuotasClientDeleteOptions) (GroupQuotasClientDeleteResponse, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, mgID, groupQuotaName, options)
	if err != nil {
		return GroupQuotasClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return GroupQuotasClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return GroupQuotasClientDeleteResponse{}, err
	}
	return GroupQuotasClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *GroupQuotasClient) deleteCreateRequest(ctx context.Context, mgID string, groupQuotaName string, options *GroupQuotasClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Management/managementGroups/{mgId}/providers/Microsoft.Quota/groupQuotas/{groupQuotaName}"
	if mgID == "" {
		return nil, errors.New("parameter mgID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{mgId}", url.PathEscape(mgID))
	if groupQuotaName == "" {
		return nil, errors.New("parameter groupQuotaName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupQuotaName}", url.PathEscape(groupQuotaName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the GroupQuotas for the name passed. It will return the GroupQuotas properties only. The details on groupQuota
// can be access from the groupQuota APIs.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01-preview
//   - mgID - Management Group Id.
//   - groupQuotaName - The GroupQuota name. The name should be unique for the provided context tenantId/MgId.
//   - options - GroupQuotasClientGetOptions contains the optional parameters for the GroupQuotasClient.Get method.
func (client *GroupQuotasClient) Get(ctx context.Context, mgID string, groupQuotaName string, options *GroupQuotasClientGetOptions) (GroupQuotasClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, mgID, groupQuotaName, options)
	if err != nil {
		return GroupQuotasClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return GroupQuotasClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return GroupQuotasClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *GroupQuotasClient) getCreateRequest(ctx context.Context, mgID string, groupQuotaName string, options *GroupQuotasClientGetOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Management/managementGroups/{mgId}/providers/Microsoft.Quota/groupQuotas/{groupQuotaName}"
	if mgID == "" {
		return nil, errors.New("parameter mgID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{mgId}", url.PathEscape(mgID))
	if groupQuotaName == "" {
		return nil, errors.New("parameter groupQuotaName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupQuotaName}", url.PathEscape(groupQuotaName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *GroupQuotasClient) getHandleResponse(resp *http.Response) (GroupQuotasClientGetResponse, error) {
	result := GroupQuotasClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GroupQuotasEntity); err != nil {
		return GroupQuotasClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists GroupQuotas for the scope passed. It will return the GroupQuotas QuotaEntity properties only.The details
// on groupQuota can be access from the groupQuota APIs.
//
// Generated from API version 2023-06-01-preview
//   - mgID - Management Group Id.
//   - options - GroupQuotasClientListOptions contains the optional parameters for the GroupQuotasClient.NewListPager method.
func (client *GroupQuotasClient) NewListPager(mgID string, options *GroupQuotasClientListOptions) *runtime.Pager[GroupQuotasClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[GroupQuotasClientListResponse]{
		More: func(page GroupQuotasClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *GroupQuotasClientListResponse) (GroupQuotasClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, mgID, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return GroupQuotasClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return GroupQuotasClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return GroupQuotasClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *GroupQuotasClient) listCreateRequest(ctx context.Context, mgID string, options *GroupQuotasClientListOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Management/managementGroups/{mgId}/providers/Microsoft.Quota/groupQuotas"
	if mgID == "" {
		return nil, errors.New("parameter mgID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{mgId}", url.PathEscape(mgID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *GroupQuotasClient) listHandleResponse(resp *http.Response) (GroupQuotasClientListResponse, error) {
	result := GroupQuotasClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GroupQuotaList); err != nil {
		return GroupQuotasClientListResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Updates the GroupQuotas for the name passed. A GroupQuotas RequestId will be returned by the Service. The
// status can be polled periodically. The status Async polling is using standards defined at -
// https://github.com/Azure/azure-resource-manager-rpc/blob/master/v1.0/async-api-reference.md#asynchronous-operations. Use
// the OperationsStatus URI provided in Azure-AsyncOperation header, the duration
// will be specified in retry-after header. Once the operation gets to terminal state - Succeeded | Failed, then the URI will
// change to Get URI and full details can be checked. Any change in the filters
// will be applicable to the future quota assignments, existing quota assigned to subscriptions from the GroupQuotas remains
// unchanged.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01-preview
//   - mgID - Management Group Id.
//   - groupQuotaName - The GroupQuota name. The name should be unique for the provided context tenantId/MgId.
//   - options - GroupQuotasClientBeginUpdateOptions contains the optional parameters for the GroupQuotasClient.BeginUpdate method.
func (client *GroupQuotasClient) BeginUpdate(ctx context.Context, mgID string, groupQuotaName string, options *GroupQuotasClientBeginUpdateOptions) (*runtime.Poller[GroupQuotasClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, mgID, groupQuotaName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[GroupQuotasClientUpdateResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[GroupQuotasClientUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Update - Updates the GroupQuotas for the name passed. A GroupQuotas RequestId will be returned by the Service. The status
// can be polled periodically. The status Async polling is using standards defined at -
// https://github.com/Azure/azure-resource-manager-rpc/blob/master/v1.0/async-api-reference.md#asynchronous-operations. Use
// the OperationsStatus URI provided in Azure-AsyncOperation header, the duration
// will be specified in retry-after header. Once the operation gets to terminal state - Succeeded | Failed, then the URI will
// change to Get URI and full details can be checked. Any change in the filters
// will be applicable to the future quota assignments, existing quota assigned to subscriptions from the GroupQuotas remains
// unchanged.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01-preview
func (client *GroupQuotasClient) update(ctx context.Context, mgID string, groupQuotaName string, options *GroupQuotasClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	req, err := client.updateCreateRequest(ctx, mgID, groupQuotaName, options)
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
func (client *GroupQuotasClient) updateCreateRequest(ctx context.Context, mgID string, groupQuotaName string, options *GroupQuotasClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Management/managementGroups/{mgId}/providers/Microsoft.Quota/groupQuotas/{groupQuotaName}"
	if mgID == "" {
		return nil, errors.New("parameter mgID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{mgId}", url.PathEscape(mgID))
	if groupQuotaName == "" {
		return nil, errors.New("parameter groupQuotaName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupQuotaName}", url.PathEscape(groupQuotaName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.GroupQuotasPatchRequestBody != nil {
		if err := runtime.MarshalAsJSON(req, *options.GroupQuotasPatchRequestBody); err != nil {
			return nil, err
		}
		return req, nil
	}
	return req, nil
}
