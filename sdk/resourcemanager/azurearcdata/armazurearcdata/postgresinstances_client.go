//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armazurearcdata

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

// PostgresInstancesClient contains the methods for the PostgresInstances group.
// Don't use this type directly, use NewPostgresInstancesClient() instead.
type PostgresInstancesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewPostgresInstancesClient creates a new instance of PostgresInstancesClient with the specified values.
//   - subscriptionID - The ID of the Azure subscription
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewPostgresInstancesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*PostgresInstancesClient, error) {
	cl, err := arm.NewClient(moduleName+".PostgresInstancesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &PostgresInstancesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreate - Creates or replaces a postgres Instance resource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-16-preview
//   - resourceGroupName - The name of the Azure resource group
//   - postgresInstanceName - Name of Postgres Instance
//   - resource - The postgres instance
//   - options - PostgresInstancesClientBeginCreateOptions contains the optional parameters for the PostgresInstancesClient.BeginCreate
//     method.
func (client *PostgresInstancesClient) BeginCreate(ctx context.Context, resourceGroupName string, postgresInstanceName string, resource PostgresInstance, options *PostgresInstancesClientBeginCreateOptions) (*runtime.Poller[PostgresInstancesClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, postgresInstanceName, resource, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[PostgresInstancesClientCreateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[PostgresInstancesClientCreateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Create - Creates or replaces a postgres Instance resource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-16-preview
func (client *PostgresInstancesClient) create(ctx context.Context, resourceGroupName string, postgresInstanceName string, resource PostgresInstance, options *PostgresInstancesClientBeginCreateOptions) (*http.Response, error) {
	var err error
	req, err := client.createCreateRequest(ctx, resourceGroupName, postgresInstanceName, resource, options)
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

// createCreateRequest creates the Create request.
func (client *PostgresInstancesClient) createCreateRequest(ctx context.Context, resourceGroupName string, postgresInstanceName string, resource PostgresInstance, options *PostgresInstancesClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureArcData/postgresInstances/{postgresInstanceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if postgresInstanceName == "" {
		return nil, errors.New("parameter postgresInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{postgresInstanceName}", url.PathEscape(postgresInstanceName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-16-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, resource); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes a postgres Instance resource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-16-preview
//   - resourceGroupName - The name of the Azure resource group
//   - postgresInstanceName - Name of Postgres Instance
//   - options - PostgresInstancesClientBeginDeleteOptions contains the optional parameters for the PostgresInstancesClient.BeginDelete
//     method.
func (client *PostgresInstancesClient) BeginDelete(ctx context.Context, resourceGroupName string, postgresInstanceName string, options *PostgresInstancesClientBeginDeleteOptions) (*runtime.Poller[PostgresInstancesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, postgresInstanceName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[PostgresInstancesClientDeleteResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[PostgresInstancesClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Deletes a postgres Instance resource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-16-preview
func (client *PostgresInstancesClient) deleteOperation(ctx context.Context, resourceGroupName string, postgresInstanceName string, options *PostgresInstancesClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, postgresInstanceName, options)
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
func (client *PostgresInstancesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, postgresInstanceName string, options *PostgresInstancesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureArcData/postgresInstances/{postgresInstanceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if postgresInstanceName == "" {
		return nil, errors.New("parameter postgresInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{postgresInstanceName}", url.PathEscape(postgresInstanceName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-16-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Retrieves a postgres Instance resource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-16-preview
//   - resourceGroupName - The name of the Azure resource group
//   - postgresInstanceName - Name of Postgres Instance
//   - options - PostgresInstancesClientGetOptions contains the optional parameters for the PostgresInstancesClient.Get method.
func (client *PostgresInstancesClient) Get(ctx context.Context, resourceGroupName string, postgresInstanceName string, options *PostgresInstancesClientGetOptions) (PostgresInstancesClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, postgresInstanceName, options)
	if err != nil {
		return PostgresInstancesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PostgresInstancesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PostgresInstancesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *PostgresInstancesClient) getCreateRequest(ctx context.Context, resourceGroupName string, postgresInstanceName string, options *PostgresInstancesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureArcData/postgresInstances/{postgresInstanceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if postgresInstanceName == "" {
		return nil, errors.New("parameter postgresInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{postgresInstanceName}", url.PathEscape(postgresInstanceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-16-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PostgresInstancesClient) getHandleResponse(resp *http.Response) (PostgresInstancesClientGetResponse, error) {
	result := PostgresInstancesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PostgresInstance); err != nil {
		return PostgresInstancesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List postgres Instance resources in the subscription
//
// Generated from API version 2023-05-16-preview
//   - options - PostgresInstancesClientListOptions contains the optional parameters for the PostgresInstancesClient.NewListPager
//     method.
func (client *PostgresInstancesClient) NewListPager(options *PostgresInstancesClientListOptions) *runtime.Pager[PostgresInstancesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[PostgresInstancesClientListResponse]{
		More: func(page PostgresInstancesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *PostgresInstancesClientListResponse) (PostgresInstancesClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return PostgresInstancesClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return PostgresInstancesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return PostgresInstancesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *PostgresInstancesClient) listCreateRequest(ctx context.Context, options *PostgresInstancesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.AzureArcData/postgresInstances"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-16-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *PostgresInstancesClient) listHandleResponse(resp *http.Response) (PostgresInstancesClientListResponse, error) {
	result := PostgresInstancesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PostgresInstanceListResult); err != nil {
		return PostgresInstancesClientListResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Get a postgres Instances list by Resource group name.
//
// Generated from API version 2023-05-16-preview
//   - resourceGroupName - The name of the Azure resource group
//   - options - PostgresInstancesClientListByResourceGroupOptions contains the optional parameters for the PostgresInstancesClient.NewListByResourceGroupPager
//     method.
func (client *PostgresInstancesClient) NewListByResourceGroupPager(resourceGroupName string, options *PostgresInstancesClientListByResourceGroupOptions) *runtime.Pager[PostgresInstancesClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[PostgresInstancesClientListByResourceGroupResponse]{
		More: func(page PostgresInstancesClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *PostgresInstancesClientListByResourceGroupResponse) (PostgresInstancesClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return PostgresInstancesClientListByResourceGroupResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return PostgresInstancesClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return PostgresInstancesClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *PostgresInstancesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *PostgresInstancesClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureArcData/postgresInstances"
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
	reqQP.Set("api-version", "2023-05-16-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *PostgresInstancesClient) listByResourceGroupHandleResponse(resp *http.Response) (PostgresInstancesClientListByResourceGroupResponse, error) {
	result := PostgresInstancesClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PostgresInstanceListResult); err != nil {
		return PostgresInstancesClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// Update - Updates a postgres Instance resource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-16-preview
//   - resourceGroupName - The name of the Azure resource group
//   - postgresInstanceName - Name of Postgres Instance
//   - parameters - The Postgres Instance.
//   - options - PostgresInstancesClientUpdateOptions contains the optional parameters for the PostgresInstancesClient.Update
//     method.
func (client *PostgresInstancesClient) Update(ctx context.Context, resourceGroupName string, postgresInstanceName string, parameters PostgresInstanceUpdate, options *PostgresInstancesClientUpdateOptions) (PostgresInstancesClientUpdateResponse, error) {
	var err error
	req, err := client.updateCreateRequest(ctx, resourceGroupName, postgresInstanceName, parameters, options)
	if err != nil {
		return PostgresInstancesClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PostgresInstancesClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PostgresInstancesClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *PostgresInstancesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, postgresInstanceName string, parameters PostgresInstanceUpdate, options *PostgresInstancesClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureArcData/postgresInstances/{postgresInstanceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if postgresInstanceName == "" {
		return nil, errors.New("parameter postgresInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{postgresInstanceName}", url.PathEscape(postgresInstanceName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-16-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *PostgresInstancesClient) updateHandleResponse(resp *http.Response) (PostgresInstancesClientUpdateResponse, error) {
	result := PostgresInstancesClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PostgresInstance); err != nil {
		return PostgresInstancesClientUpdateResponse{}, err
	}
	return result, nil
}
