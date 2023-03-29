//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

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

// OnlineDeploymentsClient contains the methods for the OnlineDeployments group.
// Don't use this type directly, use NewOnlineDeploymentsClient() instead.
type OnlineDeploymentsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewOnlineDeploymentsClient creates a new instance of OnlineDeploymentsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewOnlineDeploymentsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*OnlineDeploymentsClient, error) {
	cl, err := arm.NewClient(moduleName+".OnlineDeploymentsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &OnlineDeploymentsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create or update Inference Endpoint Deployment (asynchronous).
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - Name of Azure Machine Learning workspace.
//   - endpointName - Inference endpoint name.
//   - deploymentName - Inference Endpoint Deployment name.
//   - body - Inference Endpoint entity to apply during operation.
//   - options - OnlineDeploymentsClientBeginCreateOrUpdateOptions contains the optional parameters for the OnlineDeploymentsClient.BeginCreateOrUpdate
//     method.
func (client *OnlineDeploymentsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, endpointName string, deploymentName string, body OnlineDeployment, options *OnlineDeploymentsClientBeginCreateOrUpdateOptions) (*runtime.Poller[OnlineDeploymentsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, workspaceName, endpointName, deploymentName, body, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[OnlineDeploymentsClientCreateOrUpdateResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[OnlineDeploymentsClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Create or update Inference Endpoint Deployment (asynchronous).
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01
func (client *OnlineDeploymentsClient) createOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, endpointName string, deploymentName string, body OnlineDeployment, options *OnlineDeploymentsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, workspaceName, endpointName, deploymentName, body, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *OnlineDeploymentsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, endpointName string, deploymentName string, body OnlineDeployment, options *OnlineDeploymentsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/onlineEndpoints/{endpointName}/deployments/{deploymentName}"
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
	if endpointName == "" {
		return nil, errors.New("parameter endpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{endpointName}", url.PathEscape(endpointName))
	if deploymentName == "" {
		return nil, errors.New("parameter deploymentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deploymentName}", url.PathEscape(deploymentName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, body)
}

// BeginDelete - Delete Inference Endpoint Deployment (asynchronous).
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - Name of Azure Machine Learning workspace.
//   - endpointName - Inference endpoint name.
//   - deploymentName - Inference Endpoint Deployment name.
//   - options - OnlineDeploymentsClientBeginDeleteOptions contains the optional parameters for the OnlineDeploymentsClient.BeginDelete
//     method.
func (client *OnlineDeploymentsClient) BeginDelete(ctx context.Context, resourceGroupName string, workspaceName string, endpointName string, deploymentName string, options *OnlineDeploymentsClientBeginDeleteOptions) (*runtime.Poller[OnlineDeploymentsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, workspaceName, endpointName, deploymentName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[OnlineDeploymentsClientDeleteResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[OnlineDeploymentsClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Delete Inference Endpoint Deployment (asynchronous).
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01
func (client *OnlineDeploymentsClient) deleteOperation(ctx context.Context, resourceGroupName string, workspaceName string, endpointName string, deploymentName string, options *OnlineDeploymentsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, workspaceName, endpointName, deploymentName, options)
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
func (client *OnlineDeploymentsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, endpointName string, deploymentName string, options *OnlineDeploymentsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/onlineEndpoints/{endpointName}/deployments/{deploymentName}"
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
	if endpointName == "" {
		return nil, errors.New("parameter endpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{endpointName}", url.PathEscape(endpointName))
	if deploymentName == "" {
		return nil, errors.New("parameter deploymentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deploymentName}", url.PathEscape(deploymentName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get Inference Deployment Deployment.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - Name of Azure Machine Learning workspace.
//   - endpointName - Inference endpoint name.
//   - deploymentName - Inference Endpoint Deployment name.
//   - options - OnlineDeploymentsClientGetOptions contains the optional parameters for the OnlineDeploymentsClient.Get method.
func (client *OnlineDeploymentsClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, endpointName string, deploymentName string, options *OnlineDeploymentsClientGetOptions) (OnlineDeploymentsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, endpointName, deploymentName, options)
	if err != nil {
		return OnlineDeploymentsClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OnlineDeploymentsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return OnlineDeploymentsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *OnlineDeploymentsClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, endpointName string, deploymentName string, options *OnlineDeploymentsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/onlineEndpoints/{endpointName}/deployments/{deploymentName}"
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
	if endpointName == "" {
		return nil, errors.New("parameter endpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{endpointName}", url.PathEscape(endpointName))
	if deploymentName == "" {
		return nil, errors.New("parameter deploymentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deploymentName}", url.PathEscape(deploymentName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *OnlineDeploymentsClient) getHandleResponse(resp *http.Response) (OnlineDeploymentsClientGetResponse, error) {
	result := OnlineDeploymentsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OnlineDeployment); err != nil {
		return OnlineDeploymentsClientGetResponse{}, err
	}
	return result, nil
}

// GetLogs - Polls an Endpoint operation.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - Name of Azure Machine Learning workspace.
//   - endpointName - Inference endpoint name.
//   - deploymentName - The name and identifier for the endpoint.
//   - body - The request containing parameters for retrieving logs.
//   - options - OnlineDeploymentsClientGetLogsOptions contains the optional parameters for the OnlineDeploymentsClient.GetLogs
//     method.
func (client *OnlineDeploymentsClient) GetLogs(ctx context.Context, resourceGroupName string, workspaceName string, endpointName string, deploymentName string, body DeploymentLogsRequest, options *OnlineDeploymentsClientGetLogsOptions) (OnlineDeploymentsClientGetLogsResponse, error) {
	req, err := client.getLogsCreateRequest(ctx, resourceGroupName, workspaceName, endpointName, deploymentName, body, options)
	if err != nil {
		return OnlineDeploymentsClientGetLogsResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OnlineDeploymentsClientGetLogsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return OnlineDeploymentsClientGetLogsResponse{}, runtime.NewResponseError(resp)
	}
	return client.getLogsHandleResponse(resp)
}

// getLogsCreateRequest creates the GetLogs request.
func (client *OnlineDeploymentsClient) getLogsCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, endpointName string, deploymentName string, body DeploymentLogsRequest, options *OnlineDeploymentsClientGetLogsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/onlineEndpoints/{endpointName}/deployments/{deploymentName}/getLogs"
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
	if endpointName == "" {
		return nil, errors.New("parameter endpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{endpointName}", url.PathEscape(endpointName))
	if deploymentName == "" {
		return nil, errors.New("parameter deploymentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deploymentName}", url.PathEscape(deploymentName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, body)
}

// getLogsHandleResponse handles the GetLogs response.
func (client *OnlineDeploymentsClient) getLogsHandleResponse(resp *http.Response) (OnlineDeploymentsClientGetLogsResponse, error) {
	result := OnlineDeploymentsClientGetLogsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DeploymentLogs); err != nil {
		return OnlineDeploymentsClientGetLogsResponse{}, err
	}
	return result, nil
}

// NewListPager - List Inference Endpoint Deployments.
//
// Generated from API version 2023-04-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - Name of Azure Machine Learning workspace.
//   - endpointName - Inference endpoint name.
//   - options - OnlineDeploymentsClientListOptions contains the optional parameters for the OnlineDeploymentsClient.NewListPager
//     method.
func (client *OnlineDeploymentsClient) NewListPager(resourceGroupName string, workspaceName string, endpointName string, options *OnlineDeploymentsClientListOptions) *runtime.Pager[OnlineDeploymentsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[OnlineDeploymentsClientListResponse]{
		More: func(page OnlineDeploymentsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *OnlineDeploymentsClientListResponse) (OnlineDeploymentsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, workspaceName, endpointName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return OnlineDeploymentsClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return OnlineDeploymentsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return OnlineDeploymentsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *OnlineDeploymentsClient) listCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, endpointName string, options *OnlineDeploymentsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/onlineEndpoints/{endpointName}/deployments"
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
	if endpointName == "" {
		return nil, errors.New("parameter endpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{endpointName}", url.PathEscape(endpointName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-01")
	if options != nil && options.OrderBy != nil {
		reqQP.Set("$orderBy", *options.OrderBy)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", *options.Skip)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *OnlineDeploymentsClient) listHandleResponse(resp *http.Response) (OnlineDeploymentsClientListResponse, error) {
	result := OnlineDeploymentsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OnlineDeploymentTrackedResourceArmPaginatedResult); err != nil {
		return OnlineDeploymentsClientListResponse{}, err
	}
	return result, nil
}

// NewListSKUsPager - List Inference Endpoint Deployment Skus.
//
// Generated from API version 2023-04-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - Name of Azure Machine Learning workspace.
//   - endpointName - Inference endpoint name.
//   - deploymentName - Inference Endpoint Deployment name.
//   - options - OnlineDeploymentsClientListSKUsOptions contains the optional parameters for the OnlineDeploymentsClient.NewListSKUsPager
//     method.
func (client *OnlineDeploymentsClient) NewListSKUsPager(resourceGroupName string, workspaceName string, endpointName string, deploymentName string, options *OnlineDeploymentsClientListSKUsOptions) *runtime.Pager[OnlineDeploymentsClientListSKUsResponse] {
	return runtime.NewPager(runtime.PagingHandler[OnlineDeploymentsClientListSKUsResponse]{
		More: func(page OnlineDeploymentsClientListSKUsResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *OnlineDeploymentsClientListSKUsResponse) (OnlineDeploymentsClientListSKUsResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listSKUsCreateRequest(ctx, resourceGroupName, workspaceName, endpointName, deploymentName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return OnlineDeploymentsClientListSKUsResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return OnlineDeploymentsClientListSKUsResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return OnlineDeploymentsClientListSKUsResponse{}, runtime.NewResponseError(resp)
			}
			return client.listSKUsHandleResponse(resp)
		},
	})
}

// listSKUsCreateRequest creates the ListSKUs request.
func (client *OnlineDeploymentsClient) listSKUsCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, endpointName string, deploymentName string, options *OnlineDeploymentsClientListSKUsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/onlineEndpoints/{endpointName}/deployments/{deploymentName}/skus"
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
	if endpointName == "" {
		return nil, errors.New("parameter endpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{endpointName}", url.PathEscape(endpointName))
	if deploymentName == "" {
		return nil, errors.New("parameter deploymentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deploymentName}", url.PathEscape(deploymentName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-01")
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
func (client *OnlineDeploymentsClient) listSKUsHandleResponse(resp *http.Response) (OnlineDeploymentsClientListSKUsResponse, error) {
	result := OnlineDeploymentsClientListSKUsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SKUResourceArmPaginatedResult); err != nil {
		return OnlineDeploymentsClientListSKUsResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Update Online Deployment (asynchronous).
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - Name of Azure Machine Learning workspace.
//   - endpointName - Online Endpoint name.
//   - deploymentName - Inference Endpoint Deployment name.
//   - body - Online Endpoint entity to apply during operation.
//   - options - OnlineDeploymentsClientBeginUpdateOptions contains the optional parameters for the OnlineDeploymentsClient.BeginUpdate
//     method.
func (client *OnlineDeploymentsClient) BeginUpdate(ctx context.Context, resourceGroupName string, workspaceName string, endpointName string, deploymentName string, body PartialMinimalTrackedResourceWithSKU, options *OnlineDeploymentsClientBeginUpdateOptions) (*runtime.Poller[OnlineDeploymentsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, workspaceName, endpointName, deploymentName, body, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[OnlineDeploymentsClientUpdateResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[OnlineDeploymentsClientUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Update - Update Online Deployment (asynchronous).
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01
func (client *OnlineDeploymentsClient) update(ctx context.Context, resourceGroupName string, workspaceName string, endpointName string, deploymentName string, body PartialMinimalTrackedResourceWithSKU, options *OnlineDeploymentsClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, workspaceName, endpointName, deploymentName, body, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *OnlineDeploymentsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, endpointName string, deploymentName string, body PartialMinimalTrackedResourceWithSKU, options *OnlineDeploymentsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/onlineEndpoints/{endpointName}/deployments/{deploymentName}"
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
	if endpointName == "" {
		return nil, errors.New("parameter endpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{endpointName}", url.PathEscape(endpointName))
	if deploymentName == "" {
		return nil, errors.New("parameter deploymentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deploymentName}", url.PathEscape(deploymentName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, body)
}
