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
	"strings"
)

// EndpointDeploymentClient contains the methods for the EndpointDeployment group.
// Don't use this type directly, use NewEndpointDeploymentClient() instead.
type EndpointDeploymentClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewEndpointDeploymentClient creates a new instance of EndpointDeploymentClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewEndpointDeploymentClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*EndpointDeploymentClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &EndpointDeploymentClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create or update endpoint deployment resource with the specified parameters
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - Azure Machine Learning Workspace Name
//   - endpointName - Name of the endpoint resource.
//   - deploymentName - Name of the deployment resource
//   - body - deployment object
//   - options - EndpointDeploymentClientBeginCreateOrUpdateOptions contains the optional parameters for the EndpointDeploymentClient.BeginCreateOrUpdate
//     method.
func (client *EndpointDeploymentClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, endpointName string, deploymentName string, body EndpointDeploymentResourcePropertiesBasicResource, options *EndpointDeploymentClientBeginCreateOrUpdateOptions) (*runtime.Poller[EndpointDeploymentClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, workspaceName, endpointName, deploymentName, body, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[EndpointDeploymentClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[EndpointDeploymentClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Create or update endpoint deployment resource with the specified parameters
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-01-01-preview
func (client *EndpointDeploymentClient) createOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, endpointName string, deploymentName string, body EndpointDeploymentResourcePropertiesBasicResource, options *EndpointDeploymentClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "EndpointDeploymentClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, workspaceName, endpointName, deploymentName, body, options)
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

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *EndpointDeploymentClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, endpointName string, deploymentName string, body EndpointDeploymentResourcePropertiesBasicResource, _ *EndpointDeploymentClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/endpoints/{endpointName}/deployments/{deploymentName}"
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
	reqQP.Set("api-version", "2025-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Delete endpoint deployment resource by name
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - Azure Machine Learning Workspace Name
//   - endpointName - Name of the endpoint resource.
//   - deploymentName - Name of the deployment resource
//   - options - EndpointDeploymentClientBeginDeleteOptions contains the optional parameters for the EndpointDeploymentClient.BeginDelete
//     method.
func (client *EndpointDeploymentClient) BeginDelete(ctx context.Context, resourceGroupName string, workspaceName string, endpointName string, deploymentName string, options *EndpointDeploymentClientBeginDeleteOptions) (*runtime.Poller[EndpointDeploymentClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, workspaceName, endpointName, deploymentName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[EndpointDeploymentClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[EndpointDeploymentClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Delete endpoint deployment resource by name
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-01-01-preview
func (client *EndpointDeploymentClient) deleteOperation(ctx context.Context, resourceGroupName string, workspaceName string, endpointName string, deploymentName string, options *EndpointDeploymentClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "EndpointDeploymentClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, workspaceName, endpointName, deploymentName, options)
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
func (client *EndpointDeploymentClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, endpointName string, deploymentName string, options *EndpointDeploymentClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/endpoints/{endpointName}/deployments/{deploymentName}"
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
	reqQP.Set("api-version", "2025-01-01-preview")
	if options != nil && options.ProxyAPIVersion != nil {
		reqQP.Set("proxy-api-version", *options.ProxyAPIVersion)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get deployments under endpoint resource by name
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - Azure Machine Learning Workspace Name
//   - endpointName - Name of the endpoint resource.
//   - deploymentName - Name of the deployment resource
//   - options - EndpointDeploymentClientGetOptions contains the optional parameters for the EndpointDeploymentClient.Get method.
func (client *EndpointDeploymentClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, endpointName string, deploymentName string, options *EndpointDeploymentClientGetOptions) (EndpointDeploymentClientGetResponse, error) {
	var err error
	const operationName = "EndpointDeploymentClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, endpointName, deploymentName, options)
	if err != nil {
		return EndpointDeploymentClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return EndpointDeploymentClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return EndpointDeploymentClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *EndpointDeploymentClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, endpointName string, deploymentName string, _ *EndpointDeploymentClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/endpoints/{endpointName}/deployments/{deploymentName}"
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
	reqQP.Set("api-version", "2025-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *EndpointDeploymentClient) getHandleResponse(resp *http.Response) (EndpointDeploymentClientGetResponse, error) {
	result := EndpointDeploymentClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EndpointDeploymentResourcePropertiesBasicResource); err != nil {
		return EndpointDeploymentClientGetResponse{}, err
	}
	return result, nil
}

// NewGetInWorkspacePager - Get all the deployments under the workspace scope
//
// Generated from API version 2025-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - Azure Machine Learning Workspace Name
//   - options - EndpointDeploymentClientGetInWorkspaceOptions contains the optional parameters for the EndpointDeploymentClient.NewGetInWorkspacePager
//     method.
func (client *EndpointDeploymentClient) NewGetInWorkspacePager(resourceGroupName string, workspaceName string, options *EndpointDeploymentClientGetInWorkspaceOptions) *runtime.Pager[EndpointDeploymentClientGetInWorkspaceResponse] {
	return runtime.NewPager(runtime.PagingHandler[EndpointDeploymentClientGetInWorkspaceResponse]{
		More: func(page EndpointDeploymentClientGetInWorkspaceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *EndpointDeploymentClientGetInWorkspaceResponse) (EndpointDeploymentClientGetInWorkspaceResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "EndpointDeploymentClient.NewGetInWorkspacePager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.getInWorkspaceCreateRequest(ctx, resourceGroupName, workspaceName, options)
			}, nil)
			if err != nil {
				return EndpointDeploymentClientGetInWorkspaceResponse{}, err
			}
			return client.getInWorkspaceHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// getInWorkspaceCreateRequest creates the GetInWorkspace request.
func (client *EndpointDeploymentClient) getInWorkspaceCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, options *EndpointDeploymentClientGetInWorkspaceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/deployments"
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
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", *options.Skip)
	}
	reqQP.Set("api-version", "2025-01-01-preview")
	if options != nil && options.EndpointType != nil {
		reqQP.Set("endpointType", string(*options.EndpointType))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getInWorkspaceHandleResponse handles the GetInWorkspace response.
func (client *EndpointDeploymentClient) getInWorkspaceHandleResponse(resp *http.Response) (EndpointDeploymentClientGetInWorkspaceResponse, error) {
	result := EndpointDeploymentClientGetInWorkspaceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EndpointDeploymentResourcePropertiesBasicResourceArmPaginatedResult); err != nil {
		return EndpointDeploymentClientGetInWorkspaceResponse{}, err
	}
	return result, nil
}

// NewListPager - Get all the deployments under the endpoint resource scope
//
// Generated from API version 2025-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - Azure Machine Learning Workspace Name
//   - endpointName - Name of the endpoint resource.
//   - options - EndpointDeploymentClientListOptions contains the optional parameters for the EndpointDeploymentClient.NewListPager
//     method.
func (client *EndpointDeploymentClient) NewListPager(resourceGroupName string, workspaceName string, endpointName string, options *EndpointDeploymentClientListOptions) *runtime.Pager[EndpointDeploymentClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[EndpointDeploymentClientListResponse]{
		More: func(page EndpointDeploymentClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *EndpointDeploymentClientListResponse) (EndpointDeploymentClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "EndpointDeploymentClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, workspaceName, endpointName, options)
			}, nil)
			if err != nil {
				return EndpointDeploymentClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *EndpointDeploymentClient) listCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, endpointName string, _ *EndpointDeploymentClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/endpoints/{endpointName}/deployments"
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
	reqQP.Set("api-version", "2025-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *EndpointDeploymentClient) listHandleResponse(resp *http.Response) (EndpointDeploymentClientListResponse, error) {
	result := EndpointDeploymentClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EndpointDeploymentResourcePropertiesBasicResourceArmPaginatedResult); err != nil {
		return EndpointDeploymentClientListResponse{}, err
	}
	return result, nil
}
