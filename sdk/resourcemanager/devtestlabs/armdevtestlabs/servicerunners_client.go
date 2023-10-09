//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdevtestlabs

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

// ServiceRunnersClient contains the methods for the ServiceRunners group.
// Don't use this type directly, use NewServiceRunnersClient() instead.
type ServiceRunnersClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewServiceRunnersClient creates a new instance of ServiceRunnersClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewServiceRunnersClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ServiceRunnersClient, error) {
	cl, err := arm.NewClient(moduleName+".ServiceRunnersClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ServiceRunnersClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Create or replace an existing Service runner. This operation can take a while to complete.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - labName - The name of the lab.
//   - name - The name of the service runner.
//   - serviceRunner - A container for a managed identity to execute DevTest lab services.
//   - options - ServiceRunnersClientCreateOrUpdateOptions contains the optional parameters for the ServiceRunnersClient.CreateOrUpdate
//     method.
func (client *ServiceRunnersClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, labName string, name string, serviceRunner ServiceRunner, options *ServiceRunnersClientCreateOrUpdateOptions) (ServiceRunnersClientCreateOrUpdateResponse, error) {
	var err error
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, labName, name, serviceRunner, options)
	if err != nil {
		return ServiceRunnersClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ServiceRunnersClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return ServiceRunnersClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ServiceRunnersClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, labName string, name string, serviceRunner ServiceRunner, options *ServiceRunnersClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/servicerunners/{name}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, serviceRunner); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ServiceRunnersClient) createOrUpdateHandleResponse(resp *http.Response) (ServiceRunnersClientCreateOrUpdateResponse, error) {
	result := ServiceRunnersClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServiceRunner); err != nil {
		return ServiceRunnersClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// BeginDelete - Delete service runner. This operation can take a while to complete.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - labName - The name of the lab.
//   - name - The name of the service runner.
//   - options - ServiceRunnersClientBeginDeleteOptions contains the optional parameters for the ServiceRunnersClient.BeginDelete
//     method.
func (client *ServiceRunnersClient) BeginDelete(ctx context.Context, resourceGroupName string, labName string, name string, options *ServiceRunnersClientBeginDeleteOptions) (*runtime.Poller[ServiceRunnersClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, labName, name, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ServiceRunnersClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaOriginalURI,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[ServiceRunnersClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Delete service runner. This operation can take a while to complete.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-09-01
func (client *ServiceRunnersClient) deleteOperation(ctx context.Context, resourceGroupName string, labName string, name string, options *ServiceRunnersClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, labName, name, options)
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
func (client *ServiceRunnersClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, labName string, name string, options *ServiceRunnersClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/servicerunners/{name}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get service runner.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - labName - The name of the lab.
//   - name - The name of the service runner.
//   - options - ServiceRunnersClientGetOptions contains the optional parameters for the ServiceRunnersClient.Get method.
func (client *ServiceRunnersClient) Get(ctx context.Context, resourceGroupName string, labName string, name string, options *ServiceRunnersClientGetOptions) (ServiceRunnersClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, labName, name, options)
	if err != nil {
		return ServiceRunnersClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ServiceRunnersClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ServiceRunnersClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ServiceRunnersClient) getCreateRequest(ctx context.Context, resourceGroupName string, labName string, name string, options *ServiceRunnersClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/servicerunners/{name}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ServiceRunnersClient) getHandleResponse(resp *http.Response) (ServiceRunnersClientGetResponse, error) {
	result := ServiceRunnersClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServiceRunner); err != nil {
		return ServiceRunnersClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List service runners in a given lab.
//
// Generated from API version 2021-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - labName - The name of the lab.
//   - options - ServiceRunnersClientListOptions contains the optional parameters for the ServiceRunnersClient.NewListPager method.
func (client *ServiceRunnersClient) NewListPager(resourceGroupName string, labName string, options *ServiceRunnersClientListOptions) *runtime.Pager[ServiceRunnersClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ServiceRunnersClientListResponse]{
		More: func(page ServiceRunnersClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ServiceRunnersClientListResponse) (ServiceRunnersClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, labName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ServiceRunnersClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ServiceRunnersClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ServiceRunnersClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *ServiceRunnersClient) listCreateRequest(ctx context.Context, resourceGroupName string, labName string, options *ServiceRunnersClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/servicerunners"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Orderby != nil {
		reqQP.Set("$orderby", *options.Orderby)
	}
	reqQP.Set("api-version", "2021-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ServiceRunnersClient) listHandleResponse(resp *http.Response) (ServiceRunnersClientListResponse, error) {
	result := ServiceRunnersClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServiceRunnerList); err != nil {
		return ServiceRunnersClientListResponse{}, err
	}
	return result, nil
}
