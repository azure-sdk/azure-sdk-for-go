//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armresources

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

// DeploymentOperationsClient contains the methods for the DeploymentOperations group.
// Don't use this type directly, use NewDeploymentOperationsClient() instead.
type DeploymentOperationsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewDeploymentOperationsClient creates a new instance of DeploymentOperationsClient with the specified values.
//   - subscriptionID - The Microsoft Azure subscription ID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewDeploymentOperationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DeploymentOperationsClient, error) {
	cl, err := arm.NewClient(moduleName+".DeploymentOperationsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &DeploymentOperationsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Gets a deployments operation.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - deploymentName - The name of the deployment.
//   - operationID - The ID of the operation to get.
//   - options - DeploymentOperationsClientGetOptions contains the optional parameters for the DeploymentOperationsClient.Get
//     method.
func (client *DeploymentOperationsClient) Get(ctx context.Context, resourceGroupName string, deploymentName string, operationID string, options *DeploymentOperationsClientGetOptions) (DeploymentOperationsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, deploymentName, operationID, options)
	if err != nil {
		return DeploymentOperationsClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DeploymentOperationsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DeploymentOperationsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DeploymentOperationsClient) getCreateRequest(ctx context.Context, resourceGroupName string, deploymentName string, operationID string, options *DeploymentOperationsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/deployments/{deploymentName}/operations/{operationId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if deploymentName == "" {
		return nil, errors.New("parameter deploymentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deploymentName}", url.PathEscape(deploymentName))
	if operationID == "" {
		return nil, errors.New("parameter operationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{operationId}", url.PathEscape(operationID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DeploymentOperationsClient) getHandleResponse(resp *http.Response) (DeploymentOperationsClientGetResponse, error) {
	result := DeploymentOperationsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DeploymentOperation); err != nil {
		return DeploymentOperationsClientGetResponse{}, err
	}
	return result, nil
}

// GetAtManagementGroupScope - Gets a deployments operation.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
//   - groupID - The management group ID.
//   - deploymentName - The name of the deployment.
//   - operationID - The ID of the operation to get.
//   - options - DeploymentOperationsClientGetAtManagementGroupScopeOptions contains the optional parameters for the DeploymentOperationsClient.GetAtManagementGroupScope
//     method.
func (client *DeploymentOperationsClient) GetAtManagementGroupScope(ctx context.Context, groupID string, deploymentName string, operationID string, options *DeploymentOperationsClientGetAtManagementGroupScopeOptions) (DeploymentOperationsClientGetAtManagementGroupScopeResponse, error) {
	req, err := client.getAtManagementGroupScopeCreateRequest(ctx, groupID, deploymentName, operationID, options)
	if err != nil {
		return DeploymentOperationsClientGetAtManagementGroupScopeResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DeploymentOperationsClientGetAtManagementGroupScopeResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DeploymentOperationsClientGetAtManagementGroupScopeResponse{}, runtime.NewResponseError(resp)
	}
	return client.getAtManagementGroupScopeHandleResponse(resp)
}

// getAtManagementGroupScopeCreateRequest creates the GetAtManagementGroupScope request.
func (client *DeploymentOperationsClient) getAtManagementGroupScopeCreateRequest(ctx context.Context, groupID string, deploymentName string, operationID string, options *DeploymentOperationsClientGetAtManagementGroupScopeOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Management/managementGroups/{groupId}/providers/Microsoft.Resources/deployments/{deploymentName}/operations/{operationId}"
	if groupID == "" {
		return nil, errors.New("parameter groupID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupId}", url.PathEscape(groupID))
	if deploymentName == "" {
		return nil, errors.New("parameter deploymentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deploymentName}", url.PathEscape(deploymentName))
	if operationID == "" {
		return nil, errors.New("parameter operationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{operationId}", url.PathEscape(operationID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getAtManagementGroupScopeHandleResponse handles the GetAtManagementGroupScope response.
func (client *DeploymentOperationsClient) getAtManagementGroupScopeHandleResponse(resp *http.Response) (DeploymentOperationsClientGetAtManagementGroupScopeResponse, error) {
	result := DeploymentOperationsClientGetAtManagementGroupScopeResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DeploymentOperation); err != nil {
		return DeploymentOperationsClientGetAtManagementGroupScopeResponse{}, err
	}
	return result, nil
}

// GetAtScope - Gets a deployments operation.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
//   - scope - The resource scope.
//   - deploymentName - The name of the deployment.
//   - operationID - The ID of the operation to get.
//   - options - DeploymentOperationsClientGetAtScopeOptions contains the optional parameters for the DeploymentOperationsClient.GetAtScope
//     method.
func (client *DeploymentOperationsClient) GetAtScope(ctx context.Context, scope string, deploymentName string, operationID string, options *DeploymentOperationsClientGetAtScopeOptions) (DeploymentOperationsClientGetAtScopeResponse, error) {
	req, err := client.getAtScopeCreateRequest(ctx, scope, deploymentName, operationID, options)
	if err != nil {
		return DeploymentOperationsClientGetAtScopeResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DeploymentOperationsClientGetAtScopeResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DeploymentOperationsClientGetAtScopeResponse{}, runtime.NewResponseError(resp)
	}
	return client.getAtScopeHandleResponse(resp)
}

// getAtScopeCreateRequest creates the GetAtScope request.
func (client *DeploymentOperationsClient) getAtScopeCreateRequest(ctx context.Context, scope string, deploymentName string, operationID string, options *DeploymentOperationsClientGetAtScopeOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Resources/deployments/{deploymentName}/operations/{operationId}"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if deploymentName == "" {
		return nil, errors.New("parameter deploymentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deploymentName}", url.PathEscape(deploymentName))
	if operationID == "" {
		return nil, errors.New("parameter operationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{operationId}", url.PathEscape(operationID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getAtScopeHandleResponse handles the GetAtScope response.
func (client *DeploymentOperationsClient) getAtScopeHandleResponse(resp *http.Response) (DeploymentOperationsClientGetAtScopeResponse, error) {
	result := DeploymentOperationsClientGetAtScopeResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DeploymentOperation); err != nil {
		return DeploymentOperationsClientGetAtScopeResponse{}, err
	}
	return result, nil
}

// GetAtSubscriptionScope - Gets a deployments operation.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
//   - deploymentName - The name of the deployment.
//   - operationID - The ID of the operation to get.
//   - options - DeploymentOperationsClientGetAtSubscriptionScopeOptions contains the optional parameters for the DeploymentOperationsClient.GetAtSubscriptionScope
//     method.
func (client *DeploymentOperationsClient) GetAtSubscriptionScope(ctx context.Context, deploymentName string, operationID string, options *DeploymentOperationsClientGetAtSubscriptionScopeOptions) (DeploymentOperationsClientGetAtSubscriptionScopeResponse, error) {
	req, err := client.getAtSubscriptionScopeCreateRequest(ctx, deploymentName, operationID, options)
	if err != nil {
		return DeploymentOperationsClientGetAtSubscriptionScopeResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DeploymentOperationsClientGetAtSubscriptionScopeResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DeploymentOperationsClientGetAtSubscriptionScopeResponse{}, runtime.NewResponseError(resp)
	}
	return client.getAtSubscriptionScopeHandleResponse(resp)
}

// getAtSubscriptionScopeCreateRequest creates the GetAtSubscriptionScope request.
func (client *DeploymentOperationsClient) getAtSubscriptionScopeCreateRequest(ctx context.Context, deploymentName string, operationID string, options *DeploymentOperationsClientGetAtSubscriptionScopeOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Resources/deployments/{deploymentName}/operations/{operationId}"
	if deploymentName == "" {
		return nil, errors.New("parameter deploymentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deploymentName}", url.PathEscape(deploymentName))
	if operationID == "" {
		return nil, errors.New("parameter operationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{operationId}", url.PathEscape(operationID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getAtSubscriptionScopeHandleResponse handles the GetAtSubscriptionScope response.
func (client *DeploymentOperationsClient) getAtSubscriptionScopeHandleResponse(resp *http.Response) (DeploymentOperationsClientGetAtSubscriptionScopeResponse, error) {
	result := DeploymentOperationsClientGetAtSubscriptionScopeResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DeploymentOperation); err != nil {
		return DeploymentOperationsClientGetAtSubscriptionScopeResponse{}, err
	}
	return result, nil
}

// GetAtTenantScope - Gets a deployments operation.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
//   - deploymentName - The name of the deployment.
//   - operationID - The ID of the operation to get.
//   - options - DeploymentOperationsClientGetAtTenantScopeOptions contains the optional parameters for the DeploymentOperationsClient.GetAtTenantScope
//     method.
func (client *DeploymentOperationsClient) GetAtTenantScope(ctx context.Context, deploymentName string, operationID string, options *DeploymentOperationsClientGetAtTenantScopeOptions) (DeploymentOperationsClientGetAtTenantScopeResponse, error) {
	req, err := client.getAtTenantScopeCreateRequest(ctx, deploymentName, operationID, options)
	if err != nil {
		return DeploymentOperationsClientGetAtTenantScopeResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DeploymentOperationsClientGetAtTenantScopeResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DeploymentOperationsClientGetAtTenantScopeResponse{}, runtime.NewResponseError(resp)
	}
	return client.getAtTenantScopeHandleResponse(resp)
}

// getAtTenantScopeCreateRequest creates the GetAtTenantScope request.
func (client *DeploymentOperationsClient) getAtTenantScopeCreateRequest(ctx context.Context, deploymentName string, operationID string, options *DeploymentOperationsClientGetAtTenantScopeOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Resources/deployments/{deploymentName}/operations/{operationId}"
	if deploymentName == "" {
		return nil, errors.New("parameter deploymentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deploymentName}", url.PathEscape(deploymentName))
	if operationID == "" {
		return nil, errors.New("parameter operationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{operationId}", url.PathEscape(operationID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getAtTenantScopeHandleResponse handles the GetAtTenantScope response.
func (client *DeploymentOperationsClient) getAtTenantScopeHandleResponse(resp *http.Response) (DeploymentOperationsClientGetAtTenantScopeResponse, error) {
	result := DeploymentOperationsClientGetAtTenantScopeResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DeploymentOperation); err != nil {
		return DeploymentOperationsClientGetAtTenantScopeResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets all deployments operations for a deployment.
//
// Generated from API version 2022-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - deploymentName - The name of the deployment.
//   - options - DeploymentOperationsClientListOptions contains the optional parameters for the DeploymentOperationsClient.NewListPager
//     method.
func (client *DeploymentOperationsClient) NewListPager(resourceGroupName string, deploymentName string, options *DeploymentOperationsClientListOptions) *runtime.Pager[DeploymentOperationsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[DeploymentOperationsClientListResponse]{
		More: func(page DeploymentOperationsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DeploymentOperationsClientListResponse) (DeploymentOperationsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, deploymentName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DeploymentOperationsClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return DeploymentOperationsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DeploymentOperationsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *DeploymentOperationsClient) listCreateRequest(ctx context.Context, resourceGroupName string, deploymentName string, options *DeploymentOperationsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/deployments/{deploymentName}/operations"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if deploymentName == "" {
		return nil, errors.New("parameter deploymentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deploymentName}", url.PathEscape(deploymentName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *DeploymentOperationsClient) listHandleResponse(resp *http.Response) (DeploymentOperationsClientListResponse, error) {
	result := DeploymentOperationsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DeploymentOperationsListResult); err != nil {
		return DeploymentOperationsClientListResponse{}, err
	}
	return result, nil
}

// NewListAtManagementGroupScopePager - Gets all deployments operations for a deployment.
//
// Generated from API version 2022-09-01
//   - groupID - The management group ID.
//   - deploymentName - The name of the deployment.
//   - options - DeploymentOperationsClientListAtManagementGroupScopeOptions contains the optional parameters for the DeploymentOperationsClient.NewListAtManagementGroupScopePager
//     method.
func (client *DeploymentOperationsClient) NewListAtManagementGroupScopePager(groupID string, deploymentName string, options *DeploymentOperationsClientListAtManagementGroupScopeOptions) *runtime.Pager[DeploymentOperationsClientListAtManagementGroupScopeResponse] {
	return runtime.NewPager(runtime.PagingHandler[DeploymentOperationsClientListAtManagementGroupScopeResponse]{
		More: func(page DeploymentOperationsClientListAtManagementGroupScopeResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DeploymentOperationsClientListAtManagementGroupScopeResponse) (DeploymentOperationsClientListAtManagementGroupScopeResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listAtManagementGroupScopeCreateRequest(ctx, groupID, deploymentName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DeploymentOperationsClientListAtManagementGroupScopeResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return DeploymentOperationsClientListAtManagementGroupScopeResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DeploymentOperationsClientListAtManagementGroupScopeResponse{}, runtime.NewResponseError(resp)
			}
			return client.listAtManagementGroupScopeHandleResponse(resp)
		},
	})
}

// listAtManagementGroupScopeCreateRequest creates the ListAtManagementGroupScope request.
func (client *DeploymentOperationsClient) listAtManagementGroupScopeCreateRequest(ctx context.Context, groupID string, deploymentName string, options *DeploymentOperationsClientListAtManagementGroupScopeOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Management/managementGroups/{groupId}/providers/Microsoft.Resources/deployments/{deploymentName}/operations"
	if groupID == "" {
		return nil, errors.New("parameter groupID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupId}", url.PathEscape(groupID))
	if deploymentName == "" {
		return nil, errors.New("parameter deploymentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deploymentName}", url.PathEscape(deploymentName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listAtManagementGroupScopeHandleResponse handles the ListAtManagementGroupScope response.
func (client *DeploymentOperationsClient) listAtManagementGroupScopeHandleResponse(resp *http.Response) (DeploymentOperationsClientListAtManagementGroupScopeResponse, error) {
	result := DeploymentOperationsClientListAtManagementGroupScopeResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DeploymentOperationsListResult); err != nil {
		return DeploymentOperationsClientListAtManagementGroupScopeResponse{}, err
	}
	return result, nil
}

// NewListAtScopePager - Gets all deployments operations for a deployment.
//
// Generated from API version 2022-09-01
//   - scope - The resource scope.
//   - deploymentName - The name of the deployment.
//   - options - DeploymentOperationsClientListAtScopeOptions contains the optional parameters for the DeploymentOperationsClient.NewListAtScopePager
//     method.
func (client *DeploymentOperationsClient) NewListAtScopePager(scope string, deploymentName string, options *DeploymentOperationsClientListAtScopeOptions) *runtime.Pager[DeploymentOperationsClientListAtScopeResponse] {
	return runtime.NewPager(runtime.PagingHandler[DeploymentOperationsClientListAtScopeResponse]{
		More: func(page DeploymentOperationsClientListAtScopeResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DeploymentOperationsClientListAtScopeResponse) (DeploymentOperationsClientListAtScopeResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listAtScopeCreateRequest(ctx, scope, deploymentName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DeploymentOperationsClientListAtScopeResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return DeploymentOperationsClientListAtScopeResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DeploymentOperationsClientListAtScopeResponse{}, runtime.NewResponseError(resp)
			}
			return client.listAtScopeHandleResponse(resp)
		},
	})
}

// listAtScopeCreateRequest creates the ListAtScope request.
func (client *DeploymentOperationsClient) listAtScopeCreateRequest(ctx context.Context, scope string, deploymentName string, options *DeploymentOperationsClientListAtScopeOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Resources/deployments/{deploymentName}/operations"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if deploymentName == "" {
		return nil, errors.New("parameter deploymentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deploymentName}", url.PathEscape(deploymentName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listAtScopeHandleResponse handles the ListAtScope response.
func (client *DeploymentOperationsClient) listAtScopeHandleResponse(resp *http.Response) (DeploymentOperationsClientListAtScopeResponse, error) {
	result := DeploymentOperationsClientListAtScopeResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DeploymentOperationsListResult); err != nil {
		return DeploymentOperationsClientListAtScopeResponse{}, err
	}
	return result, nil
}

// NewListAtSubscriptionScopePager - Gets all deployments operations for a deployment.
//
// Generated from API version 2022-09-01
//   - deploymentName - The name of the deployment.
//   - options - DeploymentOperationsClientListAtSubscriptionScopeOptions contains the optional parameters for the DeploymentOperationsClient.NewListAtSubscriptionScopePager
//     method.
func (client *DeploymentOperationsClient) NewListAtSubscriptionScopePager(deploymentName string, options *DeploymentOperationsClientListAtSubscriptionScopeOptions) *runtime.Pager[DeploymentOperationsClientListAtSubscriptionScopeResponse] {
	return runtime.NewPager(runtime.PagingHandler[DeploymentOperationsClientListAtSubscriptionScopeResponse]{
		More: func(page DeploymentOperationsClientListAtSubscriptionScopeResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DeploymentOperationsClientListAtSubscriptionScopeResponse) (DeploymentOperationsClientListAtSubscriptionScopeResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listAtSubscriptionScopeCreateRequest(ctx, deploymentName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DeploymentOperationsClientListAtSubscriptionScopeResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return DeploymentOperationsClientListAtSubscriptionScopeResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DeploymentOperationsClientListAtSubscriptionScopeResponse{}, runtime.NewResponseError(resp)
			}
			return client.listAtSubscriptionScopeHandleResponse(resp)
		},
	})
}

// listAtSubscriptionScopeCreateRequest creates the ListAtSubscriptionScope request.
func (client *DeploymentOperationsClient) listAtSubscriptionScopeCreateRequest(ctx context.Context, deploymentName string, options *DeploymentOperationsClientListAtSubscriptionScopeOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Resources/deployments/{deploymentName}/operations"
	if deploymentName == "" {
		return nil, errors.New("parameter deploymentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deploymentName}", url.PathEscape(deploymentName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listAtSubscriptionScopeHandleResponse handles the ListAtSubscriptionScope response.
func (client *DeploymentOperationsClient) listAtSubscriptionScopeHandleResponse(resp *http.Response) (DeploymentOperationsClientListAtSubscriptionScopeResponse, error) {
	result := DeploymentOperationsClientListAtSubscriptionScopeResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DeploymentOperationsListResult); err != nil {
		return DeploymentOperationsClientListAtSubscriptionScopeResponse{}, err
	}
	return result, nil
}

// NewListAtTenantScopePager - Gets all deployments operations for a deployment.
//
// Generated from API version 2022-09-01
//   - deploymentName - The name of the deployment.
//   - options - DeploymentOperationsClientListAtTenantScopeOptions contains the optional parameters for the DeploymentOperationsClient.NewListAtTenantScopePager
//     method.
func (client *DeploymentOperationsClient) NewListAtTenantScopePager(deploymentName string, options *DeploymentOperationsClientListAtTenantScopeOptions) *runtime.Pager[DeploymentOperationsClientListAtTenantScopeResponse] {
	return runtime.NewPager(runtime.PagingHandler[DeploymentOperationsClientListAtTenantScopeResponse]{
		More: func(page DeploymentOperationsClientListAtTenantScopeResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DeploymentOperationsClientListAtTenantScopeResponse) (DeploymentOperationsClientListAtTenantScopeResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listAtTenantScopeCreateRequest(ctx, deploymentName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DeploymentOperationsClientListAtTenantScopeResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return DeploymentOperationsClientListAtTenantScopeResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DeploymentOperationsClientListAtTenantScopeResponse{}, runtime.NewResponseError(resp)
			}
			return client.listAtTenantScopeHandleResponse(resp)
		},
	})
}

// listAtTenantScopeCreateRequest creates the ListAtTenantScope request.
func (client *DeploymentOperationsClient) listAtTenantScopeCreateRequest(ctx context.Context, deploymentName string, options *DeploymentOperationsClientListAtTenantScopeOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Resources/deployments/{deploymentName}/operations"
	if deploymentName == "" {
		return nil, errors.New("parameter deploymentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deploymentName}", url.PathEscape(deploymentName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listAtTenantScopeHandleResponse handles the ListAtTenantScope response.
func (client *DeploymentOperationsClient) listAtTenantScopeHandleResponse(resp *http.Response) (DeploymentOperationsClientListAtTenantScopeResponse, error) {
	result := DeploymentOperationsClientListAtTenantScopeResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DeploymentOperationsListResult); err != nil {
		return DeploymentOperationsClientListAtTenantScopeResponse{}, err
	}
	return result, nil
}
