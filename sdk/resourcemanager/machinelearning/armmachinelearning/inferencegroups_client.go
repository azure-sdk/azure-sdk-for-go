//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmachinelearning

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// InferenceGroupsClient contains the methods for the InferenceGroups group.
// Don't use this type directly, use NewInferenceGroupsClient() instead.
type InferenceGroupsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewInferenceGroupsClient creates a new instance of InferenceGroupsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewInferenceGroupsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*InferenceGroupsClient, error) {
	cl, err := arm.NewClient(moduleName+".InferenceGroupsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &InferenceGroupsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create or update InferenceGroup (asynchronous).
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - Name of Azure Machine Learning workspace.
//   - poolName - InferencePool name.
//   - groupName - InferenceGroup name.
//   - body - InferenceGroup entity to apply during operation.
//   - options - InferenceGroupsClientBeginCreateOrUpdateOptions contains the optional parameters for the InferenceGroupsClient.BeginCreateOrUpdate
//     method.
func (client *InferenceGroupsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, poolName string, groupName string, body InferenceGroup, options *InferenceGroupsClientBeginCreateOrUpdateOptions) (*runtime.Poller[InferenceGroupsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, workspaceName, poolName, groupName, body, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[InferenceGroupsClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaOriginalURI,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[InferenceGroupsClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Create or update InferenceGroup (asynchronous).
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-01-preview
func (client *InferenceGroupsClient) createOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, poolName string, groupName string, body InferenceGroup, options *InferenceGroupsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, workspaceName, poolName, groupName, body, options)
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
func (client *InferenceGroupsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, poolName string, groupName string, body InferenceGroup, options *InferenceGroupsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/inferencePools/{poolName}/groups/{groupName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if poolName == "" {
		return nil, errors.New("parameter poolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{poolName}", url.PathEscape(poolName))
	if groupName == "" {
		return nil, errors.New("parameter groupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupName}", url.PathEscape(groupName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Delete InferenceGroup (asynchronous).
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - Name of Azure Machine Learning workspace.
//   - poolName - InferencePool name.
//   - groupName - InferenceGroup name.
//   - options - InferenceGroupsClientBeginDeleteOptions contains the optional parameters for the InferenceGroupsClient.BeginDelete
//     method.
func (client *InferenceGroupsClient) BeginDelete(ctx context.Context, resourceGroupName string, workspaceName string, poolName string, groupName string, options *InferenceGroupsClientBeginDeleteOptions) (*runtime.Poller[InferenceGroupsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, workspaceName, poolName, groupName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[InferenceGroupsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[InferenceGroupsClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Delete InferenceGroup (asynchronous).
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-01-preview
func (client *InferenceGroupsClient) deleteOperation(ctx context.Context, resourceGroupName string, workspaceName string, poolName string, groupName string, options *InferenceGroupsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, workspaceName, poolName, groupName, options)
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
func (client *InferenceGroupsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, poolName string, groupName string, options *InferenceGroupsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/inferencePools/{poolName}/groups/{groupName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if poolName == "" {
		return nil, errors.New("parameter poolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{poolName}", url.PathEscape(poolName))
	if groupName == "" {
		return nil, errors.New("parameter groupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupName}", url.PathEscape(groupName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get InferenceGroup.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - Name of Azure Machine Learning workspace.
//   - poolName - InferencePool name.
//   - groupName - InferenceGroup name.
//   - options - InferenceGroupsClientGetOptions contains the optional parameters for the InferenceGroupsClient.Get method.
func (client *InferenceGroupsClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, poolName string, groupName string, options *InferenceGroupsClientGetOptions) (InferenceGroupsClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, poolName, groupName, options)
	if err != nil {
		return InferenceGroupsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return InferenceGroupsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return InferenceGroupsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *InferenceGroupsClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, poolName string, groupName string, options *InferenceGroupsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/inferencePools/{poolName}/groups/{groupName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if poolName == "" {
		return nil, errors.New("parameter poolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{poolName}", url.PathEscape(poolName))
	if groupName == "" {
		return nil, errors.New("parameter groupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupName}", url.PathEscape(groupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *InferenceGroupsClient) getHandleResponse(resp *http.Response) (InferenceGroupsClientGetResponse, error) {
	result := InferenceGroupsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.InferenceGroup); err != nil {
		return InferenceGroupsClientGetResponse{}, err
	}
	return result, nil
}

// GetStatus - Retrieve inference group status.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - Name of Azure Machine Learning workspace.
//   - poolName - InferencePool name.
//   - groupName - InferenceGroup name.
//   - options - InferenceGroupsClientGetStatusOptions contains the optional parameters for the InferenceGroupsClient.GetStatus
//     method.
func (client *InferenceGroupsClient) GetStatus(ctx context.Context, resourceGroupName string, workspaceName string, poolName string, groupName string, options *InferenceGroupsClientGetStatusOptions) (InferenceGroupsClientGetStatusResponse, error) {
	var err error
	req, err := client.getStatusCreateRequest(ctx, resourceGroupName, workspaceName, poolName, groupName, options)
	if err != nil {
		return InferenceGroupsClientGetStatusResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return InferenceGroupsClientGetStatusResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return InferenceGroupsClientGetStatusResponse{}, err
	}
	resp, err := client.getStatusHandleResponse(httpResp)
	return resp, err
}

// getStatusCreateRequest creates the GetStatus request.
func (client *InferenceGroupsClient) getStatusCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, poolName string, groupName string, options *InferenceGroupsClientGetStatusOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/inferencePools/{poolName}/groups/{groupName}/getStatus"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if poolName == "" {
		return nil, errors.New("parameter poolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{poolName}", url.PathEscape(poolName))
	if groupName == "" {
		return nil, errors.New("parameter groupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupName}", url.PathEscape(groupName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getStatusHandleResponse handles the GetStatus response.
func (client *InferenceGroupsClient) getStatusHandleResponse(resp *http.Response) (InferenceGroupsClientGetStatusResponse, error) {
	result := InferenceGroupsClientGetStatusResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GroupStatus); err != nil {
		return InferenceGroupsClientGetStatusResponse{}, err
	}
	return result, nil
}

// NewListPager - List Inference Groups.
//
// Generated from API version 2024-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - Name of Azure Machine Learning workspace.
//   - poolName - Name of the InferencePool.
//   - options - InferenceGroupsClientListOptions contains the optional parameters for the InferenceGroupsClient.NewListPager
//     method.
func (client *InferenceGroupsClient) NewListPager(resourceGroupName string, workspaceName string, poolName string, options *InferenceGroupsClientListOptions) *runtime.Pager[InferenceGroupsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[InferenceGroupsClientListResponse]{
		More: func(page InferenceGroupsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *InferenceGroupsClientListResponse) (InferenceGroupsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, workspaceName, poolName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return InferenceGroupsClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return InferenceGroupsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return InferenceGroupsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *InferenceGroupsClient) listCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, poolName string, options *InferenceGroupsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/inferencePools/{poolName}/groups"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if poolName == "" {
		return nil, errors.New("parameter poolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{poolName}", url.PathEscape(poolName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-01-01-preview")
	if options != nil && options.Count != nil {
		reqQP.Set("count", strconv.FormatInt(int64(*options.Count), 10))
	}
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", *options.Skip)
	}
	if options != nil && options.Tags != nil {
		reqQP.Set("tags", *options.Tags)
	}
	if options != nil && options.Properties != nil {
		reqQP.Set("properties", *options.Properties)
	}
	if options != nil && options.OrderBy != nil {
		reqQP.Set("orderBy", string(*options.OrderBy))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *InferenceGroupsClient) listHandleResponse(resp *http.Response) (InferenceGroupsClientListResponse, error) {
	result := InferenceGroupsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.InferenceGroupTrackedResourceArmPaginatedResult); err != nil {
		return InferenceGroupsClientListResponse{}, err
	}
	return result, nil
}

// NewListSKUsPager - List Inference Group Skus.
//
// Generated from API version 2024-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - Name of Azure Machine Learning workspace.
//   - poolName - Inference Pool name.
//   - groupName - Inference Group name.
//   - options - InferenceGroupsClientListSKUsOptions contains the optional parameters for the InferenceGroupsClient.NewListSKUsPager
//     method.
func (client *InferenceGroupsClient) NewListSKUsPager(resourceGroupName string, workspaceName string, poolName string, groupName string, options *InferenceGroupsClientListSKUsOptions) *runtime.Pager[InferenceGroupsClientListSKUsResponse] {
	return runtime.NewPager(runtime.PagingHandler[InferenceGroupsClientListSKUsResponse]{
		More: func(page InferenceGroupsClientListSKUsResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *InferenceGroupsClientListSKUsResponse) (InferenceGroupsClientListSKUsResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listSKUsCreateRequest(ctx, resourceGroupName, workspaceName, poolName, groupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return InferenceGroupsClientListSKUsResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return InferenceGroupsClientListSKUsResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return InferenceGroupsClientListSKUsResponse{}, runtime.NewResponseError(resp)
			}
			return client.listSKUsHandleResponse(resp)
		},
	})
}

// listSKUsCreateRequest creates the ListSKUs request.
func (client *InferenceGroupsClient) listSKUsCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, poolName string, groupName string, options *InferenceGroupsClientListSKUsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/inferencePools/{poolName}/groups/{groupName}/skus"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if poolName == "" {
		return nil, errors.New("parameter poolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{poolName}", url.PathEscape(poolName))
	if groupName == "" {
		return nil, errors.New("parameter groupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupName}", url.PathEscape(groupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-01-01-preview")
	if options != nil && options.Count != nil {
		reqQP.Set("count", strconv.FormatInt(int64(*options.Count), 10))
	}
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", *options.Skip)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listSKUsHandleResponse handles the ListSKUs response.
func (client *InferenceGroupsClient) listSKUsHandleResponse(resp *http.Response) (InferenceGroupsClientListSKUsResponse, error) {
	result := InferenceGroupsClientListSKUsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SKUResourceArmPaginatedResult); err != nil {
		return InferenceGroupsClientListSKUsResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Update InferenceGroup (asynchronous).
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - Name of Azure Machine Learning workspace.
//   - poolName - InferencePool name.
//   - groupName - InferenceGroup name.
//   - body - Online Endpoint entity to apply during operation.
//   - options - InferenceGroupsClientBeginUpdateOptions contains the optional parameters for the InferenceGroupsClient.BeginUpdate
//     method.
func (client *InferenceGroupsClient) BeginUpdate(ctx context.Context, resourceGroupName string, workspaceName string, poolName string, groupName string, body PartialMinimalTrackedResourceWithSKU, options *InferenceGroupsClientBeginUpdateOptions) (*runtime.Poller[InferenceGroupsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, workspaceName, poolName, groupName, body, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[InferenceGroupsClientUpdateResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[InferenceGroupsClientUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Update - Update InferenceGroup (asynchronous).
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-01-preview
func (client *InferenceGroupsClient) update(ctx context.Context, resourceGroupName string, workspaceName string, poolName string, groupName string, body PartialMinimalTrackedResourceWithSKU, options *InferenceGroupsClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	req, err := client.updateCreateRequest(ctx, resourceGroupName, workspaceName, poolName, groupName, body, options)
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
func (client *InferenceGroupsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, poolName string, groupName string, body PartialMinimalTrackedResourceWithSKU, options *InferenceGroupsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/inferencePools/{poolName}/groups/{groupName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if poolName == "" {
		return nil, errors.New("parameter poolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{poolName}", url.PathEscape(poolName))
	if groupName == "" {
		return nil, errors.New("parameter groupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupName}", url.PathEscape(groupName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}
