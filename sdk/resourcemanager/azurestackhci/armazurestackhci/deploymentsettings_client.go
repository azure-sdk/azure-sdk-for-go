// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armazurestackhci

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

// DeploymentSettingsClient contains the methods for the DeploymentSettings group.
// Don't use this type directly, use NewDeploymentSettingsClient() instead.
type DeploymentSettingsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewDeploymentSettingsClient creates a new instance of DeploymentSettingsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewDeploymentSettingsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DeploymentSettingsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &DeploymentSettingsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create a DeploymentSetting
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-02-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterName - The name of the cluster.
//   - deploymentSettingsName - Name of Deployment Setting
//   - resource - Resource create parameters.
//   - options - DeploymentSettingsClientBeginCreateOrUpdateOptions contains the optional parameters for the DeploymentSettingsClient.BeginCreateOrUpdate
//     method.
func (client *DeploymentSettingsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, clusterName string, deploymentSettingsName string, resource DeploymentSetting, options *DeploymentSettingsClientBeginCreateOrUpdateOptions) (*runtime.Poller[DeploymentSettingsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, clusterName, deploymentSettingsName, resource, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[DeploymentSettingsClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[DeploymentSettingsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Create a DeploymentSetting
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-02-01-preview
func (client *DeploymentSettingsClient) createOrUpdate(ctx context.Context, resourceGroupName string, clusterName string, deploymentSettingsName string, resource DeploymentSetting, options *DeploymentSettingsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "DeploymentSettingsClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, clusterName, deploymentSettingsName, resource, options)
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
func (client *DeploymentSettingsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, deploymentSettingsName string, resource DeploymentSetting, _ *DeploymentSettingsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureStackHCI/clusters/{clusterName}/deploymentSettings/{deploymentSettingsName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if deploymentSettingsName == "" {
		return nil, errors.New("parameter deploymentSettingsName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deploymentSettingsName}", url.PathEscape(deploymentSettingsName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, resource); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Delete a DeploymentSetting
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-02-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterName - The name of the cluster.
//   - deploymentSettingsName - Name of Deployment Setting
//   - options - DeploymentSettingsClientBeginDeleteOptions contains the optional parameters for the DeploymentSettingsClient.BeginDelete
//     method.
func (client *DeploymentSettingsClient) BeginDelete(ctx context.Context, resourceGroupName string, clusterName string, deploymentSettingsName string, options *DeploymentSettingsClientBeginDeleteOptions) (*runtime.Poller[DeploymentSettingsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, clusterName, deploymentSettingsName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[DeploymentSettingsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[DeploymentSettingsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Delete a DeploymentSetting
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-02-01-preview
func (client *DeploymentSettingsClient) deleteOperation(ctx context.Context, resourceGroupName string, clusterName string, deploymentSettingsName string, options *DeploymentSettingsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "DeploymentSettingsClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, clusterName, deploymentSettingsName, options)
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
func (client *DeploymentSettingsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, deploymentSettingsName string, _ *DeploymentSettingsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureStackHCI/clusters/{clusterName}/deploymentSettings/{deploymentSettingsName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if deploymentSettingsName == "" {
		return nil, errors.New("parameter deploymentSettingsName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deploymentSettingsName}", url.PathEscape(deploymentSettingsName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get a DeploymentSetting
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-02-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterName - The name of the cluster.
//   - deploymentSettingsName - Name of Deployment Setting
//   - options - DeploymentSettingsClientGetOptions contains the optional parameters for the DeploymentSettingsClient.Get method.
func (client *DeploymentSettingsClient) Get(ctx context.Context, resourceGroupName string, clusterName string, deploymentSettingsName string, options *DeploymentSettingsClientGetOptions) (DeploymentSettingsClientGetResponse, error) {
	var err error
	const operationName = "DeploymentSettingsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, clusterName, deploymentSettingsName, options)
	if err != nil {
		return DeploymentSettingsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DeploymentSettingsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DeploymentSettingsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *DeploymentSettingsClient) getCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, deploymentSettingsName string, _ *DeploymentSettingsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureStackHCI/clusters/{clusterName}/deploymentSettings/{deploymentSettingsName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if deploymentSettingsName == "" {
		return nil, errors.New("parameter deploymentSettingsName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deploymentSettingsName}", url.PathEscape(deploymentSettingsName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DeploymentSettingsClient) getHandleResponse(resp *http.Response) (DeploymentSettingsClientGetResponse, error) {
	result := DeploymentSettingsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DeploymentSetting); err != nil {
		return DeploymentSettingsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByClustersPager - List DeploymentSetting resources by Clusters
//
// Generated from API version 2025-02-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterName - The name of the cluster.
//   - options - DeploymentSettingsClientListByClustersOptions contains the optional parameters for the DeploymentSettingsClient.NewListByClustersPager
//     method.
func (client *DeploymentSettingsClient) NewListByClustersPager(resourceGroupName string, clusterName string, options *DeploymentSettingsClientListByClustersOptions) *runtime.Pager[DeploymentSettingsClientListByClustersResponse] {
	return runtime.NewPager(runtime.PagingHandler[DeploymentSettingsClientListByClustersResponse]{
		More: func(page DeploymentSettingsClientListByClustersResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DeploymentSettingsClientListByClustersResponse) (DeploymentSettingsClientListByClustersResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "DeploymentSettingsClient.NewListByClustersPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByClustersCreateRequest(ctx, resourceGroupName, clusterName, options)
			}, nil)
			if err != nil {
				return DeploymentSettingsClientListByClustersResponse{}, err
			}
			return client.listByClustersHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByClustersCreateRequest creates the ListByClusters request.
func (client *DeploymentSettingsClient) listByClustersCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, _ *DeploymentSettingsClientListByClustersOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureStackHCI/clusters/{clusterName}/deploymentSettings"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByClustersHandleResponse handles the ListByClusters response.
func (client *DeploymentSettingsClient) listByClustersHandleResponse(resp *http.Response) (DeploymentSettingsClientListByClustersResponse, error) {
	result := DeploymentSettingsClientListByClustersResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DeploymentSettingListResult); err != nil {
		return DeploymentSettingsClientListByClustersResponse{}, err
	}
	return result, nil
}
