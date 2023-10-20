//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnginx

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

// DeploymentsClient contains the methods for the Deployments group.
// Don't use this type directly, use NewDeploymentsClient() instead.
type DeploymentsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewDeploymentsClient creates a new instance of DeploymentsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewDeploymentsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DeploymentsClient, error) {
	cl, err := arm.NewClient(moduleName+".DeploymentsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &DeploymentsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create or update the Nginx deployment
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-08-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - deploymentName - The name of targeted Nginx deployment
//   - options - DeploymentsClientBeginCreateOrUpdateOptions contains the optional parameters for the DeploymentsClient.BeginCreateOrUpdate
//     method.
func (client *DeploymentsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, deploymentName string, options *DeploymentsClientBeginCreateOrUpdateOptions) (*runtime.Poller[DeploymentsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, deploymentName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[DeploymentsClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[DeploymentsClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Create or update the Nginx deployment
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-08-01
func (client *DeploymentsClient) createOrUpdate(ctx context.Context, resourceGroupName string, deploymentName string, options *DeploymentsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, deploymentName, options)
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
func (client *DeploymentsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, deploymentName string, options *DeploymentsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Nginx.NginxPlus/nginxDeployments/{deploymentName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if deploymentName == "" {
		return nil, errors.New("parameter deploymentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deploymentName}", url.PathEscape(deploymentName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.Body != nil {
		if err := runtime.MarshalAsJSON(req, *options.Body); err != nil {
			return nil, err
		}
		return req, nil
	}
	return req, nil
}

// BeginDelete - Delete the Nginx deployment resource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-08-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - deploymentName - The name of targeted Nginx deployment
//   - options - DeploymentsClientBeginDeleteOptions contains the optional parameters for the DeploymentsClient.BeginDelete method.
func (client *DeploymentsClient) BeginDelete(ctx context.Context, resourceGroupName string, deploymentName string, options *DeploymentsClientBeginDeleteOptions) (*runtime.Poller[DeploymentsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, deploymentName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[DeploymentsClientDeleteResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[DeploymentsClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Delete the Nginx deployment resource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-08-01
func (client *DeploymentsClient) deleteOperation(ctx context.Context, resourceGroupName string, deploymentName string, options *DeploymentsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, deploymentName, options)
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
func (client *DeploymentsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, deploymentName string, options *DeploymentsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Nginx.NginxPlus/nginxDeployments/{deploymentName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if deploymentName == "" {
		return nil, errors.New("parameter deploymentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deploymentName}", url.PathEscape(deploymentName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get the Nginx deployment
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-08-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - deploymentName - The name of targeted Nginx deployment
//   - options - DeploymentsClientGetOptions contains the optional parameters for the DeploymentsClient.Get method.
func (client *DeploymentsClient) Get(ctx context.Context, resourceGroupName string, deploymentName string, options *DeploymentsClientGetOptions) (DeploymentsClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, deploymentName, options)
	if err != nil {
		return DeploymentsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DeploymentsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DeploymentsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *DeploymentsClient) getCreateRequest(ctx context.Context, resourceGroupName string, deploymentName string, options *DeploymentsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Nginx.NginxPlus/nginxDeployments/{deploymentName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if deploymentName == "" {
		return nil, errors.New("parameter deploymentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deploymentName}", url.PathEscape(deploymentName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DeploymentsClient) getHandleResponse(resp *http.Response) (DeploymentsClientGetResponse, error) {
	result := DeploymentsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Deployment); err != nil {
		return DeploymentsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List the Nginx deployments resources
//
// Generated from API version 2022-08-01
//   - options - DeploymentsClientListOptions contains the optional parameters for the DeploymentsClient.NewListPager method.
func (client *DeploymentsClient) NewListPager(options *DeploymentsClientListOptions) *runtime.Pager[DeploymentsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[DeploymentsClientListResponse]{
		More: func(page DeploymentsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DeploymentsClientListResponse) (DeploymentsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DeploymentsClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return DeploymentsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DeploymentsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *DeploymentsClient) listCreateRequest(ctx context.Context, options *DeploymentsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Nginx.NginxPlus/nginxDeployments"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *DeploymentsClient) listHandleResponse(resp *http.Response) (DeploymentsClientListResponse, error) {
	result := DeploymentsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DeploymentListResponse); err != nil {
		return DeploymentsClientListResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - List all Nginx deployments under the specified resource group.
//
// Generated from API version 2022-08-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - DeploymentsClientListByResourceGroupOptions contains the optional parameters for the DeploymentsClient.NewListByResourceGroupPager
//     method.
func (client *DeploymentsClient) NewListByResourceGroupPager(resourceGroupName string, options *DeploymentsClientListByResourceGroupOptions) *runtime.Pager[DeploymentsClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[DeploymentsClientListByResourceGroupResponse]{
		More: func(page DeploymentsClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DeploymentsClientListByResourceGroupResponse) (DeploymentsClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DeploymentsClientListByResourceGroupResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return DeploymentsClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DeploymentsClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *DeploymentsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *DeploymentsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Nginx.NginxPlus/nginxDeployments"
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
	reqQP.Set("api-version", "2022-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *DeploymentsClient) listByResourceGroupHandleResponse(resp *http.Response) (DeploymentsClientListByResourceGroupResponse, error) {
	result := DeploymentsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DeploymentListResponse); err != nil {
		return DeploymentsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Update the Nginx deployment
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-08-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - deploymentName - The name of targeted Nginx deployment
//   - options - DeploymentsClientBeginUpdateOptions contains the optional parameters for the DeploymentsClient.BeginUpdate method.
func (client *DeploymentsClient) BeginUpdate(ctx context.Context, resourceGroupName string, deploymentName string, options *DeploymentsClientBeginUpdateOptions) (*runtime.Poller[DeploymentsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, deploymentName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[DeploymentsClientUpdateResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[DeploymentsClientUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Update - Update the Nginx deployment
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-08-01
func (client *DeploymentsClient) update(ctx context.Context, resourceGroupName string, deploymentName string, options *DeploymentsClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	req, err := client.updateCreateRequest(ctx, resourceGroupName, deploymentName, options)
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

// updateCreateRequest creates the Update request.
func (client *DeploymentsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, deploymentName string, options *DeploymentsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Nginx.NginxPlus/nginxDeployments/{deploymentName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if deploymentName == "" {
		return nil, errors.New("parameter deploymentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deploymentName}", url.PathEscape(deploymentName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.Body != nil {
		if err := runtime.MarshalAsJSON(req, *options.Body); err != nil {
			return nil, err
		}
		return req, nil
	}
	return req, nil
}
