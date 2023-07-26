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

// FeaturesetContainersClient contains the methods for the FeaturesetContainers group.
// Don't use this type directly, use NewFeaturesetContainersClient() instead.
type FeaturesetContainersClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewFeaturesetContainersClient creates a new instance of FeaturesetContainersClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewFeaturesetContainersClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*FeaturesetContainersClient, error) {
	cl, err := arm.NewClient(moduleName+".FeaturesetContainersClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &FeaturesetContainersClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create or update container.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - Name of Azure Machine Learning workspace.
//   - name - Container name. This is case-sensitive.
//   - body - Container entity to create or update.
//   - options - FeaturesetContainersClientBeginCreateOrUpdateOptions contains the optional parameters for the FeaturesetContainersClient.BeginCreateOrUpdate
//     method.
func (client *FeaturesetContainersClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, name string, body FeaturesetContainer, options *FeaturesetContainersClientBeginCreateOrUpdateOptions) (*runtime.Poller[FeaturesetContainersClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, workspaceName, name, body, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[FeaturesetContainersClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaOriginalURI,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[FeaturesetContainersClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Create or update container.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01-preview
func (client *FeaturesetContainersClient) createOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, name string, body FeaturesetContainer, options *FeaturesetContainersClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, workspaceName, name, body, options)
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
func (client *FeaturesetContainersClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, name string, body FeaturesetContainer, options *FeaturesetContainersClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/featuresets/{name}"
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
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Delete container.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - Name of Azure Machine Learning workspace.
//   - name - Container name. This is case-sensitive.
//   - options - FeaturesetContainersClientBeginDeleteOptions contains the optional parameters for the FeaturesetContainersClient.BeginDelete
//     method.
func (client *FeaturesetContainersClient) BeginDelete(ctx context.Context, resourceGroupName string, workspaceName string, name string, options *FeaturesetContainersClientBeginDeleteOptions) (*runtime.Poller[FeaturesetContainersClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, workspaceName, name, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[FeaturesetContainersClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[FeaturesetContainersClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Delete container.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01-preview
func (client *FeaturesetContainersClient) deleteOperation(ctx context.Context, resourceGroupName string, workspaceName string, name string, options *FeaturesetContainersClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, workspaceName, name, options)
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
func (client *FeaturesetContainersClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, name string, options *FeaturesetContainersClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/featuresets/{name}"
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
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
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

// GetEntity - Get container.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - Name of Azure Machine Learning workspace.
//   - name - Container name. This is case-sensitive.
//   - options - FeaturesetContainersClientGetEntityOptions contains the optional parameters for the FeaturesetContainersClient.GetEntity
//     method.
func (client *FeaturesetContainersClient) GetEntity(ctx context.Context, resourceGroupName string, workspaceName string, name string, options *FeaturesetContainersClientGetEntityOptions) (FeaturesetContainersClientGetEntityResponse, error) {
	var err error
	req, err := client.getEntityCreateRequest(ctx, resourceGroupName, workspaceName, name, options)
	if err != nil {
		return FeaturesetContainersClientGetEntityResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return FeaturesetContainersClientGetEntityResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return FeaturesetContainersClientGetEntityResponse{}, err
	}
	resp, err := client.getEntityHandleResponse(httpResp)
	return resp, err
}

// getEntityCreateRequest creates the GetEntity request.
func (client *FeaturesetContainersClient) getEntityCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, name string, options *FeaturesetContainersClientGetEntityOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/featuresets/{name}"
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
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
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

// getEntityHandleResponse handles the GetEntity response.
func (client *FeaturesetContainersClient) getEntityHandleResponse(resp *http.Response) (FeaturesetContainersClientGetEntityResponse, error) {
	result := FeaturesetContainersClientGetEntityResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.FeaturesetContainer); err != nil {
		return FeaturesetContainersClientGetEntityResponse{}, err
	}
	return result, nil
}

// NewListPager - List featurestore entity containers.
//
// Generated from API version 2023-06-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - Name of Azure Machine Learning workspace.
//   - options - FeaturesetContainersClientListOptions contains the optional parameters for the FeaturesetContainersClient.NewListPager
//     method.
func (client *FeaturesetContainersClient) NewListPager(resourceGroupName string, workspaceName string, options *FeaturesetContainersClientListOptions) *runtime.Pager[FeaturesetContainersClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[FeaturesetContainersClientListResponse]{
		More: func(page FeaturesetContainersClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *FeaturesetContainersClientListResponse) (FeaturesetContainersClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, workspaceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return FeaturesetContainersClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return FeaturesetContainersClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return FeaturesetContainersClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *FeaturesetContainersClient) listCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, options *FeaturesetContainersClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/featuresets"
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01-preview")
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", *options.Skip)
	}
	if options != nil && options.Tags != nil {
		reqQP.Set("tags", *options.Tags)
	}
	if options != nil && options.ListViewType != nil {
		reqQP.Set("listViewType", string(*options.ListViewType))
	}
	if options != nil && options.PageSize != nil {
		reqQP.Set("pageSize", strconv.FormatInt(int64(*options.PageSize), 10))
	}
	if options != nil && options.Name != nil {
		reqQP.Set("name", *options.Name)
	}
	if options != nil && options.Description != nil {
		reqQP.Set("description", *options.Description)
	}
	if options != nil && options.CreatedBy != nil {
		reqQP.Set("createdBy", *options.CreatedBy)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *FeaturesetContainersClient) listHandleResponse(resp *http.Response) (FeaturesetContainersClientListResponse, error) {
	result := FeaturesetContainersClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.FeaturesetContainerResourceArmPaginatedResult); err != nil {
		return FeaturesetContainersClientListResponse{}, err
	}
	return result, nil
}
